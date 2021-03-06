package contract

import (
	"errors"
	"fmt"
	"tezosign/models"
	"tezosign/types"

	"blockwatch.cc/tzindex/micheline"
)

func BuildFullTxPayload(payload types.Payload, signatures []types.Signature) (resp []byte, entrypoint string, err error) {

	rawPayload, err := payload.MarshalBinary()
	if err != nil {
		return resp, entrypoint, err
	}

	if rawPayload[0] == TextWatermark {
		rawPayload = rawPayload[1:]
	}

	michelsonPayload := &micheline.Prim{}
	err = michelsonPayload.UnmarshalBinary(rawPayload)
	if err != nil {
		return resp, entrypoint, err
	}

	if michelsonPayload.OpCode != micheline.D_PAIR || len(michelsonPayload.Args) != 2 {
		return nil, entrypoint, fmt.Errorf("Wrong michelson payload")
	}

	signaturesParam := make([]*micheline.Prim, len(signatures))
	for i := range signatures {
		if signatures[i].IsEmpty() {
			signaturesParam[i] = &micheline.Prim{
				Type:   micheline.PrimNullary,
				OpCode: micheline.D_NONE,
			}
			continue
		}

		marshaledSig, err := signatures[i].MarshalBinary()
		if err != nil {
			return resp, entrypoint, err
		}
		signaturesParam[i] = &micheline.Prim{
			Type:   micheline.PrimUnary,
			OpCode: micheline.D_SOME,
			Args: []*micheline.Prim{
				{
					Type:   micheline.PrimBytes,
					OpCode: micheline.T_BYTES,
					//Remove address byte
					Bytes: marshaledSig,
				},
			},
		}
	}

	actionParams := &micheline.Prim{
		Type:   micheline.PrimBinary,
		OpCode: micheline.D_PAIR,
		Args: []*micheline.Prim{
			//Take params without network consts
			michelsonPayload.Args[1],
			{
				Type:   micheline.PrimSequence,
				OpCode: micheline.T_LIST,
				Args:   signaturesParam,
			},
		},
	}

	resp, err = actionParams.MarshalJSON()
	if err != nil {
		return resp, entrypoint, err
	}

	return resp, MainEntrypoint, nil
}

func GetOperationCounter(operation Operation) (counter int64, isReject bool, err error) {

	if operation.Value.OpCode != micheline.D_PAIR {
		return counter, isReject, errors.New("Wrong input param")
	}

	if operation.Value.Args[0].OpCode != micheline.D_PAIR {
		return counter, isReject, errors.New("Wrong operation param")
	}

	counter = operation.Value.Args[0].Args[0].Int.Int64()

	customPayload, err := getMichelsonParamsByActionType(models.CustomPayload, operation.Value.Args[0].Args[1])
	if err == nil {
		rejectOperation, err := customPayload.MarshalJSON()
		if err != nil {
			return counter, isReject, err
		}

		isReject = string(rejectOperation) == emptyOperation
	}

	return counter, isReject, nil
}

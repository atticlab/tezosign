package types

import (
	"encoding/hex"
	"testing"

	"blockwatch.cc/tzindex/micheline"
)

func TestTZKTPrim_UnmarshalBinary(t *testing.T) {
	type args struct {
		tzkt string
		hex  string
	}

	testCases := []struct {
		name      string
		args      args
		expResult string
		wantErr   bool
	}{
		{
			name: "Storage",
			args: args{
				tzkt: "a0070105a0070103633f210086e0a56668788415cc426351dfbb387164cd9366f9f76e87824a4e8115c2b3903f21009a3c622cb845c95c5c68800e4e1b80d0dbaec82055a33ad871103f8590c773dd3f210013b3977385b5be4c8996da7956126da11837b297823551d016df033eb8e6be9f",
			},
			expResult: `070700050707000302000000720a000000210086e0a56668788415cc426351dfbb387164cd9366f9f76e87824a4e8115c2b3900a00000021009a3c622cb845c95c5c68800e4e1b80d0dbaec82055a33ad871103f8590c773dd0a000000210013b3977385b5be4c8996da7956126da11837b297823551d016df033eb8e6be9f`,
			wantErr:   false,
		},
		{
			name: "Code param",
			args: args{
				tzkt: "9000a064816c4764656661756c74a165a065816287636f756e746572a064a164a164a164a065816e82746f816a8576616c75659163805d8a64656c65676174696f6e8d6469726563745f616374696f6ea165806ea065806ea065806e80628a7472616e73666572464186616374696f6ea05e806c905f806d87616374696f6e73a0658162897468726573686f6c64915f805c846b657973915f9063806784736967734e6d61696e5f706172616d65746572",
			},
			expResult: `05000764046c000000082564656661756c74086507650462000000083a636f756e74657207640864086408640765046e000000033a746f046a000000063a76616c75650663035d0000000b3a64656c65676174696f6e0000000e3a6469726563745f616374696f6e0865036e0765036e0765036e03620000000b3a7472616e736665724641000000073a616374696f6e075e036c055f036d000000083a616374696f6e73076504620000000a3a7468726573686f6c64065f035c000000053a6b657973065f05630367000000053a736967730000000f256d61696e5f706172616d65746572`,
			wantErr:   false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			bt, err := hex.DecodeString(test.args.tzkt)
			if err != nil {
				t.Error(err)
				return
			}

			g := &TZKTPrim{}

			err = g.UnmarshalBinary(bt)
			if err != nil {
				t.Error(err)
				return
			}

			b := micheline.Prim(*g)

			bt, err = b.MarshalBinary()
			if err != nil {
				t.Error(err)
				return
			}

			result := hex.EncodeToString(bt)

			if result != test.expResult {
				t.Errorf("wantErr: %t | results %s == %s", test.wantErr, result, test.expResult)
			}
		})
	}
}
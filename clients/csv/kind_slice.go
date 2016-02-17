package csv

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/attic-labs/noms/types"
)

type KindSlice []types.NomsKind

func (ks KindSlice) MarshalJSON() ([]byte, error) {
	elems := make([]string, len(ks))
	for i, k := range ks {
		elems[i] = fmt.Sprintf("%d", k)
	}
	return []byte("[" + strings.Join(elems, ",") + "]"), nil
}

func (ks *KindSlice) UnmarshalJSON(value []byte) error {
	elems := strings.Split(string(value[1:len(value)-1]), ",")
	*ks = make(KindSlice, len(elems))
	for i, e := range elems {
		ival, err := strconv.ParseUint(e, 10, 8)
		if err != nil {
			return err
		}
		(*ks)[i] = types.NomsKind(ival)
	}
	return nil
}

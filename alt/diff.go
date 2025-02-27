// Copyright (c) 2021, Peter Ohler, All rights reserved.

package alt

import (
	"reflect"
	"time"

	"github.com/ohler55/ojg/gen"
)

var TimeTolerance = time.Millisecond

// Path is a list of keys that can be either a string, int, or nil. Strings
// are used for keys in a map, ints are for indexes to a slice/array, and nil
// is a wildcard that matches either.
type Path []interface{}

// Diff returns the paths to the differences between two values. Any ignore
// paths are ignored in the comparison.
func Diff(v0, v1 interface{}, ignores ...Path) (diffs []Path) {
	return diff(v0, v1, false, ignores...)
}

// Compare returns a path to the first difference encountered between two
// values. Any ignore paths are ignored in the comparison.
func Compare(v0, v1 interface{}, ignores ...Path) Path {
	if diffs := diff(v0, v1, true, ignores...); 0 < len(diffs) {
		return diffs[0]
	}
	return nil
}

func diff(v0, v1 interface{}, one bool, ignores ...Path) (diffs []Path) {
	switch t0 := v0.(type) {
	case nil:
		if v1 != nil {
			diffs = append(diffs, Path{nil})
		}
	case bool:
		if t1, ok := v1.(bool); !ok || t0 != t1 {
			diffs = append(diffs, Path{nil})
		}
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		i0, _ := asInt(v0)
		if i1, ok := asInt(v1); !ok || i0 != i1 {
			diffs = append(diffs, Path{nil})
		}
	case float32, float64:
		f0, _ := asFloat(v0)
		if f1, ok := asFloat(v1); !ok || f0 != f1 {
			diffs = append(diffs, Path{nil})
		}
	case string:
		if t1, ok := v1.(string); !ok || t0 != t1 {
			diffs = append(diffs, Path{nil})
		}
	case time.Time:
		if t1, ok := v1.(time.Time); !ok || !t0.Round(TimeTolerance).Equal(t1.Round(TimeTolerance)) {
			diffs = append(diffs, Path{nil})
		}
	case []interface{}:
		t1, ok := v1.([]interface{})
		if !ok {
			diffs = append(diffs, Path{nil})
			break
		}
		var childIgnores []Path
		for _, ign := range ignores {
			if 1 < len(ign) {
				switch ign[0].(type) {
				case nil, int:
					childIgnores = append(childIgnores, ign[1:])
				}
			}
		}
		for i, m1 := range t0 {
			if ignoreIndex(i, ignores) {
				continue
			}
			if len(t1) <= i {
				diffs = append(diffs, Path{i})
				return
			}
			ds := diff(m1, t1[i], one, childIgnores...)
			for _, d := range ds {
				if len(d) == 1 && d[0] == nil {
					d[0] = i
				} else {
					d = append(Path{i}, d...)
				}
				diffs = append(diffs, d)
				if one {
					return
				}
			}
		}
		if len(t0) != len(t1) && !ignoreIndex(len(t0), ignores) {
			diffs = append(diffs, Path{len(t0)})
		}
	case map[string]interface{}:
		t1, ok := v1.(map[string]interface{})
		if !ok {
			diffs = append(diffs, Path{nil})
			break
		}
		var childIgnores []Path
		for _, ign := range ignores {
			if 1 < len(ign) {
				switch ign[0].(type) {
				case nil, string:
					childIgnores = append(childIgnores, ign[1:])
				}
			}
		}
		keys := map[string]bool{}
		for k := range t0 {
			keys[k] = true
		}
		for k := range t1 {
			keys[k] = true
		}
		for k := range keys {
			if ignoreKey(k, ignores) {
				continue
			}
			ds := diff(t0[k], t1[k], one, childIgnores...)
			for _, d := range ds {
				if len(d) == 1 && d[0] == nil {
					d[0] = k
				} else {
					d = append(Path{k}, d...)
				}
				diffs = append(diffs, d)
				if one {
					return
				}
			}
		}
	default:
		r0 := reflect.ValueOf(v0)
		r1 := reflect.ValueOf(v1)
		if r0.Type().String() == r1.Type().String() {
			if s0, _ := v0.(Simplifier); s0 != nil {
				if s1, _ := v1.(Simplifier); s1 != nil {
					return diff(s0.Simplify(), s1.Simplify(), one, ignores...)
				}
			}
			opt := &Options{}
			v0 = reflectValue(r0, opt)
			v1 = reflectValue(r1, opt)
			if v0 != nil && v1 != nil {
				return diff(v0, v1, one, ignores...)
			}
		}
		diffs = append(diffs, Path{nil})
		return
	}
	return
}

func asInt(v interface{}) (i int64, ok bool) {
	ok = true
	switch tv := v.(type) {
	case int64:
		i = tv
	case int:
		i = int64(tv)
	case int8:
		i = int64(tv)
	case int16:
		i = int64(tv)
	case int32:
		i = int64(tv)
	case uint:
		i = int64(tv)
	case uint8:
		i = int64(tv)
	case uint16:
		i = int64(tv)
	case uint32:
		i = int64(tv)
	case uint64:
		i = int64(tv)
	case float32:
		i = int64(tv)
		if float32(int64(tv)) != tv {
			ok = false
		}
	case float64:
		i = int64(tv)
		if float64(int64(tv)) != tv {
			ok = false
		}
	case gen.Int:
		i = int64(tv)
	case gen.Float:
		i = int64(tv)
		if float64(int64(tv)) != float64(tv) {
			ok = false
		}
	default:
		ok = false
	}
	return
}

func asFloat(v interface{}) (f float64, ok bool) {
	ok = true
	switch tv := v.(type) {
	case float64:
		f = tv
	case float32:
		f = float64(tv)
	case gen.Float:
		f = float64(tv)
	case int64:
		f = float64(tv)
	case int:
		f = float64(tv)
	case int8:
		f = float64(tv)
	case int16:
		f = float64(tv)
	case int32:
		f = float64(tv)
	case uint:
		f = float64(tv)
	case uint8:
		f = float64(tv)
	case uint16:
		f = float64(tv)
	case uint32:
		f = float64(tv)
	case uint64:
		f = float64(tv)
	case gen.Int:
		f = float64(tv)
	default:
		ok = false
	}
	return
}

func ignoreIndex(i int, ignores []Path) bool {
	for _, ign := range ignores {
		if len(ign) == 1 {
			switch ii := ign[0].(type) {
			case nil: // wildcard, matches any index
				return true
			case int:
				if i == ii {
					return true
				}
			}
		}
	}
	return false
}

func ignoreKey(k string, ignores []Path) bool {
	for _, ign := range ignores {
		if len(ign) == 1 {
			switch ik := ign[0].(type) {
			case nil: // wildcard, matches any index
				return true
			case string:
				if k == ik {
					return true
				}
			}
		}
	}
	return false
}

// Code generated by 'ccgo pwd/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -nostdinc -nostdlib -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -ignore-unsupported-alignment -o pwd/pwd_linux_amd64.go -pkgname pwd -Imusl-fts/ -Imusl/arch/x86_64 -Imusl/arch/generic -Imusl/obj/src/internal -Imusl/src/include -Imusl/src/internal -Imusl/obj/include -Imusl/include', DO NOT EDIT.

package pwd

import (
	"math"
	"reflect"
	"sync/atomic"
	"unsafe"
)

var _ = math.Pi
var _ reflect.Kind
var _ atomic.Value
var _ unsafe.Pointer

const (
	FEATURES_H         = 0   // features.h:2:1:
	X_BSD_SOURCE       = 1   // features.h:15:1:
	X_FEATURES_H       = 0   // features.h:2:1:
	X_FILE_OFFSET_BITS = 64  // <builtin>:25:1:
	X_LP64             = 1   // <predefined>:312:1:
	X_PWD_H            = 0   // pwd.h:2:1:
	X_STDC_PREDEF_H    = 1   // <predefined>:174:1:
	X_XOPEN_SOURCE     = 700 // features.h:16:1:
	Linux              = 1   // <predefined>:255:1:
	Unix               = 1   // <predefined>:191:1:
)

type Ptrdiff_t = int64 /* <builtin>:3:26 */

type Size_t = uint64 /* <builtin>:9:23 */

type Wchar_t = int32 /* <builtin>:15:24 */

type X__int128_t = struct {
	Flo int64
	Fhi int64
} /* <builtin>:21:43 */ // must match modernc.org/mathutil.Int128
type X__uint128_t = struct {
	Flo uint64
	Fhi uint64
} /* <builtin>:22:44 */ // must match modernc.org/mathutil.Int128

type X__builtin_va_list = uintptr /* <builtin>:46:14 */
type X__float128 = float64        /* <builtin>:47:21 */

type Uid_t = uint32 /* alltypes.h:245:18 */

type Gid_t = uint32 /* alltypes.h:250:18 */

type Passwd = struct {
	Fpw_name   uintptr
	Fpw_passwd uintptr
	Fpw_uid    Uid_t
	Fpw_gid    Gid_t
	Fpw_gecos  uintptr
	Fpw_dir    uintptr
	Fpw_shell  uintptr
} /* pwd.h:20:1 */

var _ int8 /* gen.c:2:13: */

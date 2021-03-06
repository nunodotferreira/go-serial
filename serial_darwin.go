//
// Copyright 2014-2017 Cristian Maglie. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package serial

import "syscall"

const devFolder = "/dev"
const regexFilter = "^(cu|tty)\\..*"

const ioctlTcgetattr = syscall.TIOCGETA
const ioctlTcsetattr = syscall.TIOCSETA

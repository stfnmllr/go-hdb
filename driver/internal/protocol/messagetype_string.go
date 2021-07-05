// Code generated by "stringer -type=messageType"; DO NOT EDIT.

package protocol

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[mtNil-0]
	_ = x[mtExecuteDirect-2]
	_ = x[mtPrepare-3]
	_ = x[mtAbapStream-4]
	_ = x[mtXAStart-5]
	_ = x[mtXAJoin-6]
	_ = x[mtExecute-13]
	_ = x[mtWriteLob-16]
	_ = x[mtReadLob-17]
	_ = x[mtFindLob-18]
	_ = x[mtAuthenticate-65]
	_ = x[mtConnect-66]
	_ = x[mtCommit-67]
	_ = x[mtRollback-68]
	_ = x[mtCloseResultset-69]
	_ = x[mtDropStatementID-70]
	_ = x[mtFetchNext-71]
	_ = x[mtFetchAbsolute-72]
	_ = x[mtFetchRelative-73]
	_ = x[mtFetchFirst-74]
	_ = x[mtFetchLast-75]
	_ = x[mtDisconnect-77]
	_ = x[mtExecuteITab-78]
	_ = x[mtFetchNextITab-79]
	_ = x[mtInsertNextITab-80]
	_ = x[mtBatchPrepare-81]
	_ = x[mtDBConnectInfo-82]
	_ = x[mtXopenXAStart-83]
	_ = x[mtXopenXAEnd-84]
	_ = x[mtXopenXAPrepare-85]
	_ = x[mtXopenXACommit-86]
	_ = x[mtXopenXARollback-87]
	_ = x[mtXopenXARecover-88]
	_ = x[mtXopenXAForget-89]
}

const (
	_messageType_name_0 = "mtNil"
	_messageType_name_1 = "mtExecuteDirectmtPreparemtAbapStreammtXAStartmtXAJoin"
	_messageType_name_2 = "mtExecute"
	_messageType_name_3 = "mtWriteLobmtReadLobmtFindLob"
	_messageType_name_4 = "mtAuthenticatemtConnectmtCommitmtRollbackmtCloseResultsetmtDropStatementIDmtFetchNextmtFetchAbsolutemtFetchRelativemtFetchFirstmtFetchLast"
	_messageType_name_5 = "mtDisconnectmtExecuteITabmtFetchNextITabmtInsertNextITabmtBatchPreparemtDBConnectInfomtXopenXAStartmtXopenXAEndmtXopenXAPreparemtXopenXACommitmtXopenXARollbackmtXopenXARecovermtXopenXAForget"
)

var (
	_messageType_index_1 = [...]uint8{0, 15, 24, 36, 45, 53}
	_messageType_index_3 = [...]uint8{0, 10, 19, 28}
	_messageType_index_4 = [...]uint8{0, 14, 23, 31, 41, 57, 74, 85, 100, 115, 127, 138}
	_messageType_index_5 = [...]uint8{0, 12, 25, 40, 56, 70, 85, 99, 111, 127, 142, 159, 175, 190}
)

func (i messageType) String() string {
	switch {
	case i == 0:
		return _messageType_name_0
	case 2 <= i && i <= 6:
		i -= 2
		return _messageType_name_1[_messageType_index_1[i]:_messageType_index_1[i+1]]
	case i == 13:
		return _messageType_name_2
	case 16 <= i && i <= 18:
		i -= 16
		return _messageType_name_3[_messageType_index_3[i]:_messageType_index_3[i+1]]
	case 65 <= i && i <= 75:
		i -= 65
		return _messageType_name_4[_messageType_index_4[i]:_messageType_index_4[i+1]]
	case 77 <= i && i <= 89:
		i -= 77
		return _messageType_name_5[_messageType_index_5[i]:_messageType_index_5[i+1]]
	default:
		return "messageType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
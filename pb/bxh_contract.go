package pb

func (c *StatusChange) NotifyFlags() (bool, bool) {
	if c.CurStatus == c.PrevStatus {
		return false, false
	}

	switch c.CurStatus {
	case TransactionStatus_BEGIN:
		return false, true
	case TransactionStatus_BEGIN_FAILURE:
		return true, true
	case TransactionStatus_BEGIN_ROLLBACK:
		return true, true
	case TransactionStatus_SUCCESS:
		return true, false
	case TransactionStatus_FAILURE:
		if c.PrevStatus == TransactionStatus_BEGIN {
			return true, false
		}
		return false, false
	case TransactionStatus_ROLLBACK:
		if c.PrevStatus == TransactionStatus_BEGIN || c.PrevStatus == -1 {
			return true, false
		}
		return false, false
	}

	return false, false
}

func IsFinalStatus(status TransactionStatus) bool {
	if status == TransactionStatus_SUCCESS || status == TransactionStatus_FAILURE ||
		status == TransactionStatus_ROLLBACK {
		return true
	}
	return false
}

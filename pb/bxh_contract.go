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
		return false, false
	}

	return false, false
}

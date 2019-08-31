package chkctx

import "context"

// Check quickly determines if a context is done.
func Check(ctx context.Context) (bool, error) {
	select {
	case <-ctx.Done():
		return true, ctx.Err()
	default:
	}

	return false, ctx.Err()
}

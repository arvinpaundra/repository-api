package dashboard

import "context"

type DashboardService interface {
	Overview(ctx context.Context) (int, int, int, error)
}

package service

// import "context"

// type Service struct {
// 	logger interface{}
// }

// // CONCEPT: registration?
// // CONCEPT: hooks?

// func New(log interface{}) (*Service, error) {
// 	return &Service{}
// }

// func (srv *Service) Serve(ctx context.Context) error {
// 	var err error

// 	for {
// 		select {
// 		case <-ctx.Done():
// 			logist.Info("info")
// 			logist.Info(ctx, "info")
// 			logist.Info(srv, "info")
// 			logist.Info(srv, ctx, "info")

// 			srv.logger.Info("info")
// 			srv.logger.Info(ctx, "info")

// 			// TODO: log it
// 		}
// 	}

// 	if err != nil {
// 		srv.logger.Error("error log")
// 		return err
// 	}

// 	return nil
// }

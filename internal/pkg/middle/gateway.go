package middle

type Res struct {
	Code int
	Msg  string
}

//func ErrHandler() runtime.ErrorHandlerFunc {
//	return func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {
//		log.Println("http url = ", request.RequestURI)
//		s := Res{
//			Code: 200,
//		}
//		statusErr, ok := status.FromError(err)
//		s.Msg = err.Error()
//		if ok && int32(statusErr.Code()) == int32(basepb.ErrorCode_EC_BUSINESSERR) { //
//			s.Code = int(basepb.ErrorCode_EC_BUSINESSERR)
//			s.Msg = statusErr.Message()
//			errs := make([]zap.Field, 0, len(statusErr.Proto().GetDetails()))
//			for _, detail := range statusErr.Proto().GetDetails() {
//				errs = append(errs, zap.Error(errors.New(string(detail.Value))))
//
//			}
//			logger.Error(s.Msg, errs...)
//		}
//		bytes, err := marshaler.Marshal(s)
//		if err != nil {
//			logger.Error("err = ", zap.Error(err))
//		}
//		_, err = writer.Write(bytes)
//		if err != nil {
//			logger.Error("err = ", zap.Error(err))
//		}
//	}
//}

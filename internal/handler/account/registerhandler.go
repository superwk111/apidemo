package account

import (
	"net/http"

	"apidemo/internal/logic/account"
	"apidemo/internal/svc"
	"apidemo/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"apidemo/internal/common"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, common.InvalidParamsError())
			return
		}

		l := account.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			if codeErr, ok := common.IsCodeError(err); ok {
				httpx.OkJsonCtx(r.Context(), w, codeErr)
			} else {
				httpx.ErrorCtx(r.Context(), w, common.HandleError(err))
			}
		} else {
			httpx.OkJsonCtx(r.Context(), w, common.Success(resp))
		}
	}
}

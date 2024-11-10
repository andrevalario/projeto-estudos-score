package mdlmiddleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/andrevalario/projeto-estudos-score/domain"
	mdlusuario "github.com/andrevalario/projeto-estudos-score/model/usuario"
	"github.com/andrevalario/projeto-estudos-score/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

// Função para validar o token JWT
func ValidarToken(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			utils.ErrorResponseJson(r.Context(), w, fmt.Errorf("token não encontrado"))
			return
		}

		token, err := jwt.Parse(strings.TrimPrefix(tokenString, "Bearer "), func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("método de assinatura inválido")
			}
			return []byte(os.Getenv("JWT_TOKEN")), nil
		})

		if err != nil || !token.Valid {
			utils.ErrorResponseJson(r.Context(), w, fmt.Errorf("token inválido"))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.ErrorResponseJson(r.Context(), w, fmt.Errorf("erro ao ler as claims do token"))
			return
		}

		idUsuarioClaims, ok := claims["id"].(float64)
		if !ok {
			utils.ErrorResponseJson(r.Context(), w, fmt.Errorf("ID inválido"))
			return
		}

		idUsuario := uint64(idUsuarioClaims)
		usuario, err := mdlusuario.FetchById(idUsuario)
		if err != nil {
			return
		}

		if usuario.Id == 0 {
			utils.ErrorResponseJson(r.Context(), w, fmt.Errorf("usuário não encontrado"))
		}

		ctx := context.WithValue(r.Context(), domain.UsuarioAutenticado, &usuario)
		r = r.WithContext(ctx)

		h(w, r, ps)
	}
}

func ValidarAcessoDivida(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		usuario, existe := r.Context().Value(domain.UsuarioAutenticado).(*domain.Usuario)
		if !existe {
			utils.ErrorResponseJson(r.Context(), w, fmt.Errorf("usuário não encontrado no contexto"))
			return
		}

		if usuario.TipoUsuario != domain.Admin {
			utils.ErrorResponseJson(r.Context(), w, fmt.Errorf("acesso não permitido para este tipo de usuário"))
		}

		h(w, r, ps)
	}
}

func ValidarAcessoBens(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		usuario, existe := r.Context().Value(domain.UsuarioAutenticado).(*domain.Usuario)
		if !existe {
			utils.ErrorResponseJson(r.Context(), w, fmt.Errorf("usuário não encontrado no contexto"))
			return
		}

		// Usuário admin tem acesos total
		if usuario.TipoUsuario == domain.Admin {
			h(w, r, ps)
			return
		}

		//Validacao de pertencimento do dado levado para o usecase para não duplicar chamadas ao banco de dados.

		h(w, r, ps)
	}
}

func GetUsuarioAutenticado(ctx context.Context) (domain.Usuario, error) {
	usuarioAutenticado, existe := ctx.Value(domain.UsuarioAutenticado).(*domain.Usuario)
	if !existe {
		fmt.Println("to caindo aqui")
		return domain.Usuario{}, fmt.Errorf("usuário aiutenticado não encontrado")
	}

	return *usuarioAutenticado, nil
}

func ValidarAcessoAdmin(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		usuario, existe := r.Context().Value(domain.UsuarioAutenticado).(*domain.Usuario)
		if !existe {
			utils.ErrorResponseJson(r.Context(), w, fmt.Errorf("usuário não encontrado no contexto"))
			return
		}

		if usuario.TipoUsuario != domain.Admin {
			utils.ErrorResponseJson(r.Context(), w, fmt.Errorf("acesso não permitido para este tipo de usuário"))
		}

		h(w, r, ps)
	}
}

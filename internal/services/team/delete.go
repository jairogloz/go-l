package team

import (
	"context"
	"errors"
	"log"

	"github.com/jairogloz/go-l/internal/domain"
	"github.com/jairogloz/go-l/internal/repositories/mongo/team"
)

/*
Delete recibe un contexto y un id de tipo string, y retorna un error.

Parámetros:
	- ctx contexto
	- id string

El método llama al método Delete del campo Repo de la estructura Service, la id string y el contexto como argumentos.

Si el método Repo.Delete devuelve un error, el método Delete team service verifica el error e registra el error
y devuelve un nuevo error que envuelve el error original con un mensaje que indica que hubo un error al eliminar el equipo.

Valores de retorno:
  - err: Un error que será nil si el equipo se elimino correctamente. Si hubo un error, será un objeto de error que describe el fallo.

Nota: Dejo team.DeletedAt comentado porque no se está utilizando en el código, ya que aun no se hace el core de la funcionalidad.


*/

func (s Service) Delete(ctx context.Context, id string) (err error) {
	//now = time.Now().UTC()
	//team.DeletedAt = &now

	err = s.Repo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, team.ErrIncorrectHexID) {
			log.Println("converting hex to object")
			appErr := domain.AppError{
				Code: team.ErrIncorrectHexID.Error(),
				Msg:  "error deleting team: converting hex to object id failed",
			}
			return appErr
		}
		if errors.Is(err, team.ErrTeamNotFound) {
			log.Printf("team with id %s not found", id)
			appErr := domain.AppError{
				Code: domain.ErrCodeNotFound,
				Msg:  "error deleting team: team not found",
			}
			return appErr
		}
		if errors.Is(err, team.ErrDeleteTeam) {
			log.Println(err.Error())
			log.Printf("error deleting team with id %s", id)
			appErr := domain.AppError{
				Code: team.ErrDeleteTeam.Error(),
				Msg:  "error deleting team: error deleting team",
			}
			return appErr
		}
		return err
	}
	return nil
}

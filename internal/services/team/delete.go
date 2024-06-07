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

Si el método Repo.Delete devuelve un error, el método Delete team service verifica. Si lo es, registra el error
y devuelve un nuevo AppError con el código y mensaje de error de clave duplicada. Si el error no es un error de clave duplicada, registra el error
y devuelve un nuevo error que envuelve el error original con un mensaje que indica que hubo un error al crear el equipo.

Valores de retorno:
  - err: Un error que será nil si el equipo se creó correctamente. Si hubo un error, será un objeto de error que describe el fallo.

Nota: Dejo team.CreatedAt comentado porque no se está utilizando en el código, ya que aun no se hace el core de la funcionalidad.


*/

func (s Service) Delete(ctx context.Context, id string) (err error) {
	//now = time.Now().UTC()
	//team.DeleteAt = &now

	err = s.Repo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, team.ErrTeamNotFound) {
			log.Default().Printf("team with id %s not found", id)
			appErr := domain.AppError{
				Code: domain.ErrCodeNotFound,
				Msg:  "error deleting team: team not found",
			}
			return appErr
		}
		if errors.Is(err, team.ErrDeleteTeam) {
			log.Println(err.Error())
		return err
	}
	return nil
}

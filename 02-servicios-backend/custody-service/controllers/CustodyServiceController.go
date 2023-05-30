package controllers

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	pb "github.com/malarcon-79/microservices-lab/grpc-protos-go/system/custody"

	"github.com/malarcon-79/microservices-lab/orm-go/dao"
	"github.com/malarcon-79/microservices-lab/orm-go/model"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Controlador de servicio gRPC
type CustodyServiceController struct {
	logger *zap.SugaredLogger // Logger
	re     *regexp.Regexp     // Expresión regular para validar formato de períodos YYYY-MM
}

// Método a nivel de package, permite construir una instancia correcta del controlador de servicio gRPC
func NewCustodyServiceController() (CustodyServiceController, error) {
	_logger, _ := zap.NewProduction() // Generamos instancia de logger
	logger := _logger.Sugar()

	re, err := regexp.Compile(`^\d{4}\-(0?[1-9]|1[012])$`) // Expresión regular para validar períodos YYYY-MM
	if err != nil {
		return CustodyServiceController{}, err
	}

	instance := CustodyServiceController{
		logger: logger, // Asignamos el logger
		re:     re,     // Asignamos el RegExp precompilado
	}
	return instance, nil // Devolvemos la nueva instancia de este Struct y un puntero nulo para el error
}

func (c *CustodyServiceController) AddCustodyStock(ctx context.Context, msg *pb.CustodyAdd) (*pb.Empty, error) {
	orm := dao.DB.Model(&model.Custody{})

	// TODO: Validaciones
	print(msg)
	// Creamos el modelo de datos para almacenamiento
	custody := &model.Custody{
		Period:   msg.Period,
		ClientId: msg.ClientId,
		Stock:    msg.Stock,
		// Market:   "Market de prueba",
		// Price:    decimal.NewFromInt(2),
		Quantity: int32(msg.Quantity),
	}

	// Insert
	if err := orm.Save(custody).Error; err != nil {
		print(err)
		c.logger.Error("No se pudo agregar la custodia correctamente", err)
		return nil, errors.New("error al guardar")
	}

	print("Pasamos el supuesto insert")
	// Implementar este método
	return &pb.Empty{}, nil
	// return nil, errors.New("no implementado")
}

func (c *CustodyServiceController) ClosePeriod(ctx context.Context, msg *pb.CloseFilters) (*pb.Empty, error) {
	return nil, errors.New("no implementadoaaaa")
}

func (c *CustodyServiceController) GetCustody(ctx context.Context, msg *pb.CustodyFilter) (*pb.Custodies, error) {

	fmt.Println("GetCustody")
	// Con esta línea instanciamos el ORM para trabajar con la tabla "Custody"
	orm := dao.DB.Model(&model.Custody{})

	// Arreglo de punteros a registros de tabla "Custody"
	custodys := []*model.Custody{}
	// Creamos el filtro de búsqueda usando los campos del mismo modelo
	filter := &model.Custody{
		Period:        msg.Period,
		ClientId:      msg.ClientId,
	}
	// Ejecutamos el SELECT con un Inner Join (instrucción Preload) sobre la relación y evaluamos si hubo errores
	if err := orm.Preload("Custody").Find(&custodys, filter).Error; err != nil {
		c.logger.Errorf("no se pudo buscar facturas con filtros %v", filter, err)
		return nil, status.Errorf(codes.Internal, "no se pudo realizar query")
	}

	// Este será el mensaje de salida
	result := &pb.Custodies{}

	print("Imprimo")
	print(result.GetItems)



	// Implementar este método
	return nil, errors.New("No implementado000000")
}

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


	//custodys := []*model.Custody{}
	// TODO: Validaciones
	print(msg)

	//period: período de la custodia en formato YYYY-MM. Parte de la llave primaria (PK) del registro de custodia. No puede ser nulo
    //stock: nemotécnico del instrumento en custodia. Parte de la llave primaria. No puede ser nulo
	//client_id: identificador del cliente (RUT). Parte de la llave primaria. No puede ser nulo
	//quantity: cantidad de instrumentos en custodia. Debe ser mayor o igual a cero

	// Validacion del periodo
	if msg.Period == "" {
		c.logger.Error("Periodo no puede estar en blanco")
		return nil, errors.New("Periodo no puede estar en blanco")
	}

	// Validacion del nemo
	if msg.Stock == "" {
		c.logger.Error("Stock no puede estar en blanco")
		return nil, errors.New("Stock no puede estar en blanco")
	}

	// Validacion del rut
	if msg.ClientId == "" {
		c.logger.Error("Rut no puede estar en blanco")
		return nil, errors.New("Rut no puede estar en blanco")
	}

	if msg.Quantity <= 0 {
		c.logger.Error("Cantidad no puede ser menor o igual a 0")
		return nil, errors.New("Cantidad no puede ser menor o igual a 0")

	}

	// Arreglo de punteros a registros de tabla "Custody"
	custodys := model.Custody{}
	// Creamos el filtro de búsqueda usando los campos del mismo modelo
	filter := &model.Custody{
		Period:   msg.Period,
		Stock:    msg.Stock,
		ClientId: msg.ClientId,
	}
	if err := orm.First(&custodys, filter).Error; err != nil {
		c.logger.Errorf("No se pudo encontrar la custodia con filtros %v", filter, err)
		return nil, status.Errorf(codes.Internal, "No se pudo encontrar la custodia")
	}


	print(custodys.Quantity)

	// Creamos el modelo de datos para almacenamiento
	custody := &model.Custody{
		Period:   msg.Period,
		ClientId: msg.ClientId,
		Stock:    msg.Stock,
		Quantity: int32(msg.Quantity) + int32(custodys.Quantity),
	}

	// Insert
	if err := orm.Save(custody).Error; err != nil {
		print(err)
		c.logger.Error("No se pudo agregar la custodia correctamente", err)
		return nil, errors.New("error al guardar")
	}

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
		Period:   msg.Period,
		Stock:    msg.Stock,
		ClientId: msg.ClientId,
	}
	// Ejecutamos el SELECT con un Inner Join (instrucción Preload) sobre la relación y evaluamos si hubo errores
	if err := orm.Find(&custodys, filter).Error; err != nil {
		c.logger.Errorf("no se pudo buscar facturas con filtros %v", filter, err)
		return nil, status.Errorf(codes.Internal, "no se pudo realizar query")
	}

	// Este será el mensaje de salida
	result := &pb.Custodies{}

	print("Imprimo")
	print(result.GetItems)

	for _, item := range custodys {

		result.Items = append(result.Items, &pb.Custodies_Custody {
			Period:        item.Period,
			Stock:         item.Stock,
			ClientId:      item.ClientId,
			Market:        item.Market,
			Price:    	   item.Price.InexactFloat64(), // Pasamos de decimal.Decimal a Float64
			Quantity:      item.Quantity,
		})
	}

	// Implementar este método
	return result, nil

}

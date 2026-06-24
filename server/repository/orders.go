package repository

import (
	"database/sql"
	"log"

	"github.com/VelVit24/projext/dto"
	"github.com/VelVit24/projext/models"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) InsertOrder(id_user int, order *models.Order) error {
	cart, err := r.SelectCart(id_user)
	if err != nil || cart != nil || len(cart) == 0 {
		return sql.ErrNoRows
	}
	var user models.User
	err = r.db.QueryRow("select phone, email, first_name, last_name from users where id=$1", id_user).Scan(&user.Phone, &user.Email, &user.FirstName, &user.LastName)
	if err != nil {
		return err
	}
	err = r.db.QueryRow("insert into orders(id_user, phone, email, first_name, last_name) values ($1, $2, $3, $4, $5) returning id", id_user, user.Phone, user.Email, user.FirstName, user.LastName).Scan(&order.Id)
	if err != nil {
		return err
	}
	carts, err := r.SelectCart(id_user)
	total := 0
	for _, cart := range carts {
		total += cart.Price * cart.Amount
		_, err := r.db.Exec("insert into order_items(id_order, id_product, amount, price) values ($1, $2, $3, $4)", order.Id, cart.Id_product, cart.Amount, cart.Price)
		if err != nil {
			log.Println(err)
		}
	}
	_, err = r.db.Exec("update orders set total=$1 where id=$2", total, order.Id)
	if err != nil {
		return err
	}
	_, err = r.db.Exec("delete from cart_items where id_user=$1", id_user)
	return err
}

func (r *OrderRepository) UpdateOrder(order *models.Order) error {
	res, err := r.db.Exec("update orders set status=$1 where id=$2", order.Status, order.Id)
	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}
	return err
}

func (r *OrderRepository) DeleteOrder(id int) error {
	res, err := r.db.Exec("delete from order where id=$1", id)
	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}
	if err != nil {
		return err
	}
	res, err = r.db.Exec("delete from order_items where id_order=$1", id)
	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}
	return err
}

func (r *OrderRepository) SelectOrders(id_user, page, limit int, role string) (dto.OrderResponce, error) {
	responce := dto.OrderResponce{}
	rows, err := r.db.Query("select o.id, status, total, created_at, phone, email, first_name, last_name from orders where id_user=$1 offset $2 limit $3", id_user, (page-1)*limit, limit)
	if err != nil {
		return dto.OrderResponce{}, err
	}
	for rows.Next() {
		order := dto.OrderFull{}
		err := rows.Scan(&order.Order.Id, &order.Order.Status, &order.Order.Total, &order.Order.CreatedAt, &order.Order.Phone, &order.Order.Email, &order.Order.FirstName, &order.Order.LastName)
		if err != nil {
			log.Println(err)
		}
		items, err := r.db.Query("select i.id_product, p.name, i.amount, i.price from order_items i left outer join products p on p.id=i.id_product where id_order=$1", order.Order.Id)
		if err != nil {
			log.Println(err)
			responce.Orders = append(responce.Orders, order)
			continue
		}
		for items.Next() {
			item := dto.CartItems{}
			err := items.Scan(&item.Id_product, &item.Name, &item.Amount, &item.Price)
			if err != nil {
				log.Println(err)
			}
			order.OrderItems = append(order.OrderItems, item)
		}
		responce.Orders = append(responce.Orders, order)
	}
	responce.Page = page
	responce.Limit = limit
	err = r.db.QueryRow("select count(*) from orders where id_user=$1", id_user).Scan(&responce.Total)
	if err != nil {
		return dto.OrderResponce{}, err
	}
	return responce, nil
}

func (r *OrderRepository) SelectOrdersAdmin(id_user, page, limit int, role string) (dto.OrderResponce, error) {
	responce := dto.OrderResponce{}
	rows, err := r.db.Query("select o.id, o.id_user, status, total, created_at, phone, email, first_name, last_name from orders o left outer join users u on o.id_user=u.id offset $1 limit $2", (page-1)*limit, limit)
	if err != nil {
		return dto.OrderResponce{}, err
	}
	for rows.Next() {
		order := dto.OrderFull{}
		id_user := 0
		err := rows.Scan(&order.Order.Id, &id_user, &order.Order.Status, &order.Order.Total, &order.Order.CreatedAt, &order.Order.Phone, &order.Order.Email, &order.Order.FirstName, &order.Order.LastName)
		if err != nil {
			log.Println(err)
		}
		err = r.db.QueryRow("select id, email, first_name, last_name from users where id=$1", id_user).Scan(&order.User.Id, &order.User.Email, &order.User.FirstName, &order.User.LastName)
		if err != nil {
			log.Println(err)
		}
		items, err := r.db.Query("select i.id_product, p.name, i.amount, i.price from order_items i left outer join products p on p.id=i.id_product where id_order=$1", order.Order.Id)
		if err != nil {
			log.Println(err)
			responce.Orders = append(responce.Orders, order)
			continue
		}
		for items.Next() {
			item := dto.CartItems{}
			err := items.Scan(&item.Id_product, &item.Name, &item.Amount, &item.Price)
			if err != nil {
				log.Println(err)
			}
			order.OrderItems = append(order.OrderItems, item)
		}
		responce.Orders = append(responce.Orders, order)
	}
	responce.Page = page
	responce.Limit = limit
	err = r.db.QueryRow("select count(*) from orders where id_user=$1", id_user).Scan(&responce.Total)
	if err != nil {
		return dto.OrderResponce{}, err
	}
	return responce, nil
}

func (r *OrderRepository) SelectCart(id_user int) ([]dto.CartItems, error) {
	rows, err := r.db.Query("select id_product, name, price, stock, image_count, amount from cart_items c left outer join products p on c.id_product = p.id where id_user=$1", id_user)
	if err != nil {
		return nil, err
	}
	items := []dto.CartItems{}
	for rows.Next() {
		item := dto.CartItems{}
		err := rows.Scan(&item.Id_product, &item.Name, &item.Price, &item.Stock, &item.Image_count, &item.Amount)
		if err != nil {
			log.Println(err)
		}
		if item.Amount > item.Stock {
			item.IsInStock = false
		} else {
			item.IsInStock = true
		}
		items = append(items, item)
	}
	return items, err
}
func (r *OrderRepository) InsertOrderNoAuth(orderRequest dto.OrderRequestNoAuth) error {
	id := 0
	err := r.db.QueryRow("insert into orders(phone, email, first_name, last_name) values ($1, $2, $3, $4) returning id", orderRequest.Phone, orderRequest.Email, orderRequest.FirstName, orderRequest.LastName).Scan(&id)
	if err != nil {
		return err
	}
	total := 0
	for _, item := range orderRequest.CartItems {
		total += item.Price * item.Amount
		_, err := r.db.Exec("insert into order_items(id_order, id_product, amount, price) values ($1, $2, $3, $4)", id, item.Id_product, item.Amount, item.Price)
		if err != nil {
			log.Println(err)
		}
	}
	_, err = r.db.Exec("update orders set total=$1 where id=$2", total, id)
	if err != nil {
		return err
	}
	return nil
}

package database

import "errors"



var(

	ErrCantProduct = errors.New("can't find the product")
	ErrCantDecodeProducts  = errors.New("can't find the product")
	ErrUserIdIsNotValid    = errors.New("this user is not valid")
    ErrCantUpdateUser = errors.New("cannot add this product to the cart")
	ErrCantRemoveItemCart = errors.New("cannot remove this item from the cart")
	ErrCantGetItem   = errors.New("was unable to get the item from the cart")
	ErrCantBUyCartItem  = errors.New("cannot update the purchase")

)


func AddProductToCart(){

}

func RemoveCartItem(){

}


func BuyFromCart(){


}

func InstantBuyer(){

}
package domain

type CustomerRepositoryStub struct{
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error){
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001" ,"Samanta" ,"Cape Town", "1234", "1995-10-06", "1"},
		{"1002" ,"John" ,"New Yor", "1255", "1995-10-06", "1"},
	}

	return CustomerRepositoryStub{customers}

}
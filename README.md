## Requirements

- Golang
- Postgres
- Redis-Server

## Setup

```shell
git clone https://github.com/anujkumarthakur/uber-demo
```

```shell
go inside uber folder 
```

```shell
go get
```

```shell
database configure .env 
```

```shell
create database uber;
```

```
insert into pincode(country_code,placename,city,latitude,longitude,accuracy)
values

('IN/110007','Birla Lines','New Delhi','28.6833','77.2',4);
('IN/110008','Patel Nagar','New Delhi','28.65','77.2167',4),
('IN/110010','Delhi Cantt','New Delhi','28.55','77.2667',4);


create table users(user_id SERIAL PRIMARY KEY, 
                    name VARCHAR(200), 
                    username VARCHAR(200), 
                    password VARCHAR(200), 
                    mobile VARCHAR(10),
                );

create table usersbookinghistory(
    booking_id SERIAL,
    user_id int, 
    name VARCHAR(200), 
    username VARCHAR(200),  
    mobile VARCHAR(10),
    from_loc VARCHAR(200),
    destination VARCHAR(200),
    bookingtime VARCHAR(200),
    driver_name VARCHAR(200),
    driver_car_no VARCHAR(200),
    total_fare int,
    PRIMARY KEY(booking_id),
      FOREIGN KEY(user_id) 
	  REFERENCES users(user_id)
);


create table drivers(driver_id uuid DEFAULT uuid_generate_v4 (),
        name VARCHAR(200), 
        mobile VARCHAR(200), 
        Address VARCHAR(200),
        car_no VARCHAR(100),
        c_latitude VARCHAR(200),
        c_longitude VARCHAR(200)
    );
``` 

```shell
go build -o uber
```

```shell
./uber 
```

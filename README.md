# Food Delivery App

work in progress...

## Stack

- Backend: Go microservices + gRPC with protocol buffers
- Frontend: Vue 3 + TypeScript
- Database for each microservice: PostgreSQL

## Just to make it clear

![Project architecture design](design.png)

## Getting Started

Since there's no admin panel yet for adding restaurants and menu items, you'll need to insert them manually into the database.

### Sample Data

**Insert restaurants:**

```sql
INSERT INTO restaurants (id, name, description, address, phone, latitude, longitude, opening_time, closing_time, logo_url)
VALUES
('c796574c-4712-411d-91b4-5262c0879c94', 'KFC', 'Знаменитая жареная курица, баскеты и популярные блюда быстрого питания.', 'улица Толе би, 151', '7797', 43.252815, 76.911796, '10:00:00', '23:45:00', 'https://www.kfc.kz/admin/files/3190.svg'),
('d887685d-5823-522e-a2c5-6373d1980d05', 'dodo pizza', 'Сеть пиццерий с быстрой доставкой, пиццей и закусками.', 'улица Толе би, 100', '+7‒7719444004', 43.252796, 76.926129, '08:00:00', '03:00:00', 'https://dodopizza.kz/logo.png');
```

**Insert menu items:**

```sql
INSERT INTO menu_items (id, restaurant_id, name, description, price, image_url, category)
VALUES
(gen_random_uuid(), 'c796574c-4712-411d-91b4-5262c0879c94', 'Сандерс баскет', '1 ножка оригинальная, 2 крыла, 2 стрипса оригинальных, 2 наггетса.', 2500, 'https://www.kfc.kz/admin/files/5379.jpg', 'Баскет'),
(gen_random_uuid(), 'c796574c-4712-411d-91b4-5262c0879c94', 'Баскет 6 крыльев', '6 острых крыльев.', 2700, 'https://www.kfc.kz/admin/files/5384.jpg', 'Баскет'),
(gen_random_uuid(), 'c796574c-4712-411d-91b4-5262c0879c94', 'Шефбургер комбо Острый', 'Шефбургер, картофель фри большой, Pepsi 0,5 л, кетчуп.', 3050, 'https://www.kfc.kz/admin/files/5336.png', 'Комбо'),
(gen_random_uuid(), 'c796574c-4712-411d-91b4-5262c0879c94', 'Кранч Мастер Острый Комбо L', 'Кранч Мастер Острый, картофель фри большой, Pepsi 0,5 л, 1 соус.', 3440, 'https://www.kfc.kz/admin/files/5207.jpg', 'Комбо'),
(gen_random_uuid(), 'c796574c-4712-411d-91b4-5262c0879c94', 'Байтс & Фри сырный', 'Картофель фри, сырный соус, байтсы.', 2200, 'https://www.kfc.kz/admin/files/5315.png', 'Снэки'),
(gen_random_uuid(), 'c796574c-4712-411d-91b4-5262c0879c94', 'Pepsi 0,5 л', 'Пепси 0,5 л из холодильника.', 790, 'https://www.kfc.kz/admin/files/5324.jpg', 'Напитки'),
(gen_random_uuid(), 'd887685d-5823-522e-a2c5-6373d1980d05', 'Чикен Бомбони', 'Куриные кусочки, перец, соус сладкий чили, моцарелла, смесь сыров, лук, соус Альфредо.', 3190, 'https://media.dodostatic.net/image/r:292x292/019b2193c0b4786d8b688fddde3fa8de.avif', 'Пицца'),
(gen_random_uuid(), 'd887685d-5823-522e-a2c5-6373d1980d05', 'Четыре сезона', 'Моцарелла, ветчина из цыпленка, пепперони, брынза, томаты, шампиньоны, томатный соус, итальянские травы.', 2990, 'https://media.dodostatic.net/image/r:292x292/01995c479e6e7430b77b3b72a73d0416.avif', 'Пицца'),
(gen_random_uuid(), 'd887685d-5823-522e-a2c5-6373d1980d05', 'Диабло', 'Острая чоризо, халапеньо, соус барбекю, митболы из говядины, томаты, сладкий перец, красный лук, моцарелла.', 2650, 'https://media.dodostatic.net/image/r:292x292/01995c6b69e67950bb90f7d0cc61220c.avif', 'Пицца'),
(gen_random_uuid(), 'd887685d-5823-522e-a2c5-6373d1980d05', 'Додстер', 'Горячая закуска с цыплёнком, томатами, моцареллой и соусом ранч в тонкой лепешке.', 1890, 'https://media.dodostatic.net/image/r:292x292/0198eb2d853f768894e8b9f8e1e2f945.avif', 'Закуски'),
(gen_random_uuid(), 'd887685d-5823-522e-a2c5-6373d1980d05', 'Дэнвич ветчина и сыр', 'Чиабатта с цыплёнком, ветчиной из цыплёнка, моцареллой, томатами и соусом ранч.', 2090, 'https://media.dodostatic.net/image/r:292x292/019897d1339a74d89a6300dfc822114d.avif', 'Закуски');
```

Now you can start testing the API endpoints with real data.

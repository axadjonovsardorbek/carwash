INSERT INTO users (id, email, password, first_name, last_name, phone, role, created_at, updated_at, deleted_at) VALUES
('91cf6e44-d61a-4e81-8516-f906ef8c53e8', 'axadjonovsardorbeck@gmail.com', '$2a$10$DCHM3DqLWoA.lgdqM7Tkk.Qdq/OHMkBq5DaM6TCYpQQKmdF7tmfQW', 'Sardor', 'Axadjonov', '+998200070424', 'customer', '2024-08-13 13:39:43.540641+05', '2024-08-13 13:39:43.540641+05', 0),
('91cf6e44-d61a-4e81-8516-f906ef8c53f9', 'axadjonovsardorbekk@gmail.com', '$2a$10$DCHM3DqLWoA.lgdqM7Tkk.Qdq/OHMkBq5DaM6TCYpQQKmdF7tmfQW', 'Sardor', 'Axadjonov', '+998909082075', 'provider', '2024-08-13 13:39:43.540641+05', '2024-08-13 13:39:43.540641+05', 0),
('b2c9f3f4-9c8a-42e1-97a3-72a2f1e2a1e6', 'john.doe@example.com', '$2a$10$DCHM3DqLWoA.lgdqM7Tkk.Qdq/OHMkBq5DaM6TCYpQQKmdF7tmfQW', 'John', 'Doe', '+998901234567', 'customer', '2024-08-13 14:00:00.000000+05', '2024-08-13 14:00:00.000000+05', 0),
('d3f5a3f7-3c7c-4d1f-8092-2a3d6f7e4c1d', 'jane.smith@example.com', '$2a$10$DCHM3DqLWoA.lgdqM7Tkk.Qdq/OHMkBq5DaM6TCYpQQKmdF7tmfQW', 'Jane', 'Smith', '+998902345678', 'provider', '2024-08-13 14:15:00.000000+05', '2024-08-13 14:15:00.000000+05', 0),
('e6c9b5f4-9b7a-43e1-87b3-32a3e1e4f3f5', 'alice.wonder@example.com', '$2a$10$DCHM3DqLWoA.lgdqM7Tkk.Qdq/OHMkBq5DaM6TCYpQQKmdF7tmfQW', 'Alice', 'Wonder', '+998903456789', 'admin', '2024-08-13 14:30:00.000000+05', '2024-08-13 14:30:00.000000+05', 0);

INSERT INTO providers (id, user_id, company_name, description, availability, average_rating, location, created_at, updated_at, deleted_at)
VALUES
('47c7a789-73fe-4bba-87c8-7e75a1b642e4', '91cf6e44-d61a-4e81-8516-f906ef8c53f9', 'Car Care Inc.', 'Professional car care services', '09:00-18:00', 4.5, 'Downtown', '2024-08-13 13:39:43.540641+05', '2024-08-15 13:53:40.313415+05', 0);

INSERT INTO provider_services (id, user_id, service_id, provider_id, created_at, updated_at, deleted_at)
VALUES
('82b326b7-e316-4e0f-9f7f-4066364ec04c', '91cf6e44-d61a-4e81-8516-f906ef8c53f9', '11a9c6f4-2b7a-43e1-97a3-72a2f1e2b2e1', '47c7a789-73fe-4bba-87c8-7e75a1b642e4', '2024-08-13 13:39:43.540641+05', '2024-08-13 13:39:43.540641+05', 0),
('828d3e0f-70e9-4cf6-8a32-3686164151cd', '91cf6e44-d61a-4e81-8516-f906ef8c53f9', '22b0d7f5-3c8b-54f2-08b4-82b3f2e3c3f2', '47c7a789-73fe-4bba-87c8-7e75a1b642e4', '2024-08-13 14:00:00+05', '2024-08-13 14:00:00+05', 0),
('b6eb9431-b36d-4738-8af1-127a0a1e21ce', '91cf6e44-d61a-4e81-8516-f906ef8c53f9', '2723561c-7c0a-4974-85ec-57582aa38ebb', '47c7a789-73fe-4bba-87c8-7e75a1b642e4', '2024-08-15 14:53:18.623737+05', '2024-08-15 14:53:18.623737+05', 1723715628);

INSERT INTO bookings (id, user_id, provider_id, service_id, status, scheduled_time, location, total_price, created_at, updated_at, deleted_at)
VALUES
('c02a6b53-3309-4040-a155-0a7abcb637aa', '91cf6e44-d61a-4e81-8516-f906ef8c53e8', '47c7a789-73fe-4bba-87c8-7e75a1b642e4', '22b0d7f5-3c8b-54f2-08b4-82b3f2e3c3f2', 'completed', '2024-09-09 18:00:00+05', 'string', 75000, '2024-08-15 13:13:59.168035+05', '2024-08-15 14:06:24.217561+05', 0),
('95fe8f04-29f2-4da9-8236-da863f320410', '91cf6e44-d61a-4e81-8516-f906ef8c53e8', '47c7a789-73fe-4bba-87c8-7e75a1b642e4', '22b0d7f5-3c8b-54f2-08b4-82b3f2e3c3f2', 'cancelled', '2024-09-09 18:00:00+05', 'string', 75000, '2024-08-15 13:06:29.746293+05', '2024-08-15 13:20:08.390969+05', 0);


INSERT INTO payments (id, user_id, booking_id, amount, status, payment_method, created_at, updated_at, deleted_at)
VALUES
('2d7504ac-3650-46cc-af53-9a99d98b8617', '91cf6e44-d61a-4e81-8516-f906ef8c53e8', 'c02a6b53-3309-4040-a155-0a7abcb637aa', 75000, 'completed', 'cash', '2024-08-15 13:16:49.419058+05', '2024-08-15 14:06:24.193073+05', 0);

INSERT INTO reviews (id, booking_id, user_id, provider_id, rating, comment, created_at, updated_at, deleted_at)
VALUES
('33dc2282-00c6-4fef-adcf-7ddd512660f7', 'c02a6b53-3309-4040-a155-0a7abcb637aa', '91cf6e44-d61a-4e81-8516-f906ef8c53e8', '47c7a789-73fe-4bba-87c8-7e75a1b642e4', 4, 'boladi chidasa', '2024-08-15 13:38:50.929579+05', '2024-08-15 13:38:50.929579+05', 0),
('3294f6bd-3f10-4779-b876-ca478c3de24b', '95fe8f04-29f2-4da9-8236-da863f320410', '91cf6e44-d61a-4e81-8516-f906ef8c53e8', '47c7a789-73fe-4bba-87c8-7e75a1b642e4', 5, 'hammasi chotki', '2024-08-15 13:39:32.488195+05', '2024-08-15 13:53:40.299718+05', 0);

INSERT INTO services (id, name, description, price, duration, created_at, updated_at, deleted_at)
VALUES
('11a9c6f4-2b7a-43e1-97a3-72a2f1e2b2e1', 'Car Wash', 'Complete exterior and interior car wash', 50000, 60, '2024-08-13 13:39:43.540641+05', '2024-08-13 13:39:43.540641+05', 0),
('22b0d7f5-3c8b-54f2-08b4-82b3f2e3c3f2', 'Oil Change', 'Engine oil and filter replacement', 75000, 30, '2024-08-13 14:00:00+05', '2024-08-13 14:00:00+05', 0),
('5e4818de-81ae-45e6-9737-959cc3bfb834', 'polni', 'polni moyka', 200000, 20, '2024-08-15 14:18:50.446065+05', '2024-08-15 14:18:50.446065+05', 0),
('2723561c-7c0a-4974-85ec-57582aa38ebb', 'polni', 'polni moyka', 200000, 20, '2024-08-15 14:18:51.806645+05', '2024-08-15 14:18:51.806645+05', 0),
('2fc61f9f-06a5-4d82-bd6f-0f4e2a2cdb0a', 'full moyka', 'moyka full', 20000, 20, '2024-08-15 14:21:32.665996+05', '2024-08-15 14:21:32.665996+05', 1723714007),
('8541932e-f06a-4e3a-8625-4605892855e6', 'polni', 'polni moyka', 50000, 20, '2024-08-15 14:19:04.139478+05', '2024-08-15 14:30:22.681519+05', 0);

### setup

- [x] Go Fiber
- [x] Go Websocket
- [x] sqlite
- [x] svelte
- [x] connect

### create ws event & API route

## ws event

- [ ] **user event**
   1. `order_food`

   ```text
   "payload": {
       orderId: uuid()
       menuId: array of uuid()
       tableId: staticId
   }
   ```

   2. `request_checkout`

   ```text
   "payload": {
       tableId: staticId
   }
   ```

- [ ] **kitchen event**
   1. `update_order_status` : Prepared | complete

   ```text
   "payload": {
       tableId
       orderId
       newStatus
   }
   ```

   2. `toggle_menu_status`

   ```text
   "payload": {
       menuId
       newStatus
   }
   ```

- [ ] **server event**
   1. `send_slip_result`
   ```text
   "payload": {
       tableId
       result: status code
       detail: long txt
   }
   ```

- [ ] API route

- [ ] add ratelimit 

- [x] get menu: `GET /api/menu`
    ```text
        response
        {
            Title,
            Desc,
            Price,
            Is_available
        }
   ```
- [ ] upload slip: `POST /api/payments/slip`
    ```text
        response
        {
            table_id
            img: base64
        }
   ```


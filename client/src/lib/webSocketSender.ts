export function SendCreateOrderReq(socket: WebSocket, tableId: number) {
	if (socket && socket.readyState === WebSocket.OPEN) {
		const eventPayload = {
			type: 'create_order',
			payload: {
				tableId
			}
		};

		socket.send(JSON.stringify(eventPayload));
	}
}

export function SendOrderFoodReq(
	socket: WebSocket,
	orderId: number,
	orderMap: Map<number, number>,
	tableId: number
) {
	if (socket && socket.readyState === WebSocket.OPEN) {
		const eventPayload = {
			type: 'order_food',
			payload: {
				orderMap,
				tableId,
				orderId
			}
		};

		socket.send(JSON.stringify(eventPayload));
	}
}
export function SendRequestCheckoutReq(){
    
}

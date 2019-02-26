
let msg = {
    name: 'channel add',
    data: {
        name: 'Hardware Support',
    }
}

let subMsg = {
    name: 'channel subscribe'
}


let ws = new WebSocket('ws://192.168.99.100:3001');

ws.onopen = () => {
    ws.send(JSON.stringify(msg));
}

let username = prompt("Enter your name to join the chat");
const ws = new WebSocket("ws://localhost:8800/api/ws?username=" + username);
ws.onmessage = (e) => {
    const message = JSON.parse(e.data);
    console.log(message)

    document.querySelector(".chat-messages").innerHTML += `
        <div class="card">
            <div class="card-header">
                ${message.sender}
            </div>
            <div class="card-body">
                <p class="card-text">${message.content}</p>
            </div>
        </div>
            
    `
};

function Message(props) {
    return (

        <div class="card">
            <div class="card-header">
                {props.sender}
            </div>
            <div class="card-body">
                <p class="card-text">{props.content}</p>
            </div>
        </div>

        // <div class="message">
        //     <h4 class="meta">{props.sender}</h4>
        //     <p class="text">
        //         {props.content}
        //     </p>
        // </div>
    )
}

function SendingForm(props) {
    const sendMessage = async (e) => {
        let chatRoomId;
        e.preventDefault();
        e.stopPropagation();
        const msg = document.getElementById("msg").value;
        const to = document.getElementById("to").value;
        let chatroom = await (await fetch("http://localhost:8800/api/chat-room/members?member1=" + username + "&member2=" + to)).json()
        console.log(chatroom)
        if (chatroom.member == null) {
            let result = await
                fetch("http://localhost:8800/api/chat-room", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify({
                        name: "Chatroom" + username + to,
                        member: [username, to]
                    })
                })
            chatRoomId = (await result.json()).InsertedID;
        } else {
            chatRoomId = chatroom.id;
        }
        console.log(chatRoomId)
        const message = {
            chatRoom: chatRoomId,
            sender: username,
            content: msg,
        };
        ws.send(JSON.stringify(message));
        document.getElementById("msg").value = "";
        document.getElementById("to").value = "";
    };
    return (
        <div class="container">
            <div class="chat">
                <div class="chat-header">
                    <h1>Simple Chat</h1>
                    <h2>User: {username}</h2>
                </div>
                <div class="chat-messages"></div>
                <form >
                    <div class="form-group">
                        <label for="msg" class="form-label">Message</label>
                        <textarea class="form-control" id="msg" rows="3" type="text" placeholder="Enter Message" required autocomplete="off" />
                    </div>
                    <div class="form-group">
                        <label for="to" class="form-label">Receiver</label>
                        <input class="form-control" id="to" type="text" placeholder="To" required autocomplete="off" />
                    </div>

                    <button type="submit" class="btn btn-primary mb-3" onClick={sendMessage}>Send</button>

                </form>

            </div>
        </div>
    );
}

window.onload = async () => {
    const root = ReactDOM.createRoot(
        document.getElementById("root")
    );

    root.render(SendingForm());
}


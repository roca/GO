import React, {Component} from 'react';
import ChannelSection from './channels/ChannelSection.jsx';
import UserSection from './users/UserSection.jsx';
import MessageSection from './messages/MessageSection.jsx';

class App extends Component{
    constructor(props){
        super(props);
        this.state = {
            channels: [],
            users: [],
            messages: [],
            activeChannel: {},
            connected: false
        };
    }

    setChannel(activeChannel) {
        this.setState({activeChannel});
        //console.log('Get Channels Message');
        // TODO: Get Channels Message
    }   

    componentDidMount() {
        let ws = this.ws = new WebSocket('ws://echo.websocket.org');
        ws.onmessage = this.message.bind(this);
        ws.onopen = this.open.bind(this);
        ws.onclose = this.close.bind(this);
    }

    message(e) {
        const event = JSON.parse(e.data);
        if(event.name === 'channel add'){
            this.newChannel(event.data);
        }
    }

    open(){
        this.setState({connected: true});
    }

    close() {
        this.setState({connected: false});
    }

    newChannel(channel){
        let {channels} = this.state;
        channels.push(channel);
        this.setState({channels});
    }

    addChannel(name) {
        let {channels} = this.state;
         // TODO: Send to server
         let msg = {
             name: 'channel add',
             data: {
                 id: channels.length,
                 name 
             }
         }
         this.ws.send(JSON.stringify(msg));
    }

    setUserName(name) {
        let {users} = this.state;
        users.push({id: users.length, name});
        this.setState({users});
         // TODO: Send to server
    }

    addMessage(body) {
        let{messages, users} = this.state;
        let createdAt = new Date;
        let author = users.length > 0 ? users[0].name : 'anonymous';
        messages.push({id: messages.length, body, createdAt, author});
        this.setState({messages});
        // TODO: Send to server

    }
    render() {
        return (
            <div className='app'>
                <div className='nav'>
                    <ChannelSection 
                        {...this.state} 
                        setChannel={this.setChannel.bind(this)} 
                        addChannel={this.addChannel.bind(this)}/>
                    <UserSection 
                        {...this.state} 
                        setUserName={this.setUserName.bind(this)}/>
                </div>
                <MessageSection
                    {...this.state}
                    addMessage={this.addMessage.bind(this)}/>
            </div>
         );
    }
}


export default App;
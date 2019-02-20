import React, {Component} from 'react';
import ChannelSection from './channels/ChannelSection.jsx';
import UserSection from './users/UserSection.jsx';
import MessageSection from './messages/MessageSection.jsx';
import Socket from './../socket.js';
import { SSL_OP_COOKIE_EXCHANGE } from 'constants';

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
        let socket = this.socket = new Socket();
        socket.on('connect', this.OnConnect.bind(this));
        socket.on('disconnect', this.onDisConnect.bind(this));
        socket.on('channel add', this.onAddChannel.bind(this));
        socket.on('user add', this.onAddUser.bind(this));
        socket.on('user edit', this.onEditUser.bind(this));
        socket.on('user remove', this.onRemoveUser.bind(this));
    }

    onRemoveUser(removeUser){
        let {users} = this.state;
        users = users.filter(user => {
            return user.id !== removeUser.id;
        });
        this.setState({users});   
    }

    onAddUser(user) {
        let {users} = this.state;
        users.push(user);
        this.setState({users});
    }

    onEditUser(editUser){
        let {users} = this.state;
        users = users.map(user => {
            if(editUser.id === user.id) {
                return editUser;
            }
            return user;
        });
        this.setState({users});
    }

    onConnect(){
        this.setState({connected: true});
    }

    onDisConnect(){
        this.setState({connected: false});
    }
    newChannel(channel){
        let {channels} = this.state;
        channels.push(channel);
        this.setState({channels});
    }

    addChannel(name) {
        this.socket.emit('channel add', {name});
    }

    setUserName(name) {
        this.socket.emit('user edit', {name});
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
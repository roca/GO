import React, {Component} from 'react';
import ChannelSection from './channels/ChannelSection.jsx';
import UserSection from './users/UserSection.jsx';

class App extends Component{
    constructor(props){
        super(props);
        this.state = {
            channels: [],
            users: [],
            activeChannel: {}
        };
    }

    setChannel(activeChannel) {
        this.setState({activeChannel});
        console.log('Get Channels Message');
        // TODO: Get Channels Message
    }   

    addChannel(name) {
        let {channels} = this.state;
        channels.push({id: channels.length, name});
        this.setState({channels});
         // TODO: Send to server
    }

    setUserName(name) {
        let {users} = this.state;
        users.push({id: users.length, name});
        this.setState({users});
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
                
            </div>
         );
    }
}


export default App;
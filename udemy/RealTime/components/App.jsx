import React, {Component} from 'react';
import ChannelSection from './channels/ChannelSection.jsx';

class App extends Component{
    constructor(props){
        super(props);
        this.state = {
            channels: [],
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

    render() {
        return (
            <div className='app'>
                <div className='nav'>
                    <ChannelSection 
                        {...this.state} 
                        setChannel={this.setChannel.bind(this)} 
                        addChannel={this.addChannel.bind(this)}/>
                </div>
            </div>
         );
    }
}


export default App;
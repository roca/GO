class ChannelSection extends React.Component{
    constructor(props){
        super(props);
        this.state = {
            channels: [
                {name: 'Hardware Support'},
                {name: 'Software Support'}
            ]
        };
    }
    addChannel(name){
        let {channels} = this.state;
        channels.push({name: name});
        this.setState({
            channels: channels
        });
    }
    render() {
        return(
            <div>
                <ChannelList channels={this.state.channels} />
                <ChannelForm addChannel={this.addChannel.bind(this)}/>
            </div>
        )
    }
}
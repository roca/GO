let channels =[
    {name: 'Hardware Support'},
    {name: 'Software Support'}
];

class Channel extends React.Component{
    onClick() {
        console.log('I was clicked',this.props.name);
    }
    render() {
        return(
            <li  onClick={this.onClick.bind(this)}>{this.props.name}</li>
        )
    }
}

class ChannelList extends React.Component {
    render() {
        return (
            <ul>
               {this.props.channels.map( (channel,index) => {
                   return(
                       <Channel key={index} name={channel.name}/>
                   )
                  }
               )}
            </ul>
        )
    }
}

class ChannelForm extends React.Component {
    constructor(props){
        super(props);
        this.state = {};
    }
    onChange(e){
        this.setState({
            channelName: e.target.value 
        });
        //console.log(e.target.value);
    }
    onSubmit(e) {
        let {channelName} = this.state;
        console.log(channelName);
        channels.push({
            name: channelName
        });
        this.setState({
            channelName: ''
        });
        e.preventDefault();
    }
    render() {
        return (
            <form onSubmit={this.onSubmit.bind(this)}>
                <input type='text'
                    onChange={this.onChange.bind(this)}
                    value={this.state.channelName}></input>
            </form>
        )
    }
}

class ChannelSection extends React.Component{
    render() {
        return(
            <div>
                <ChannelList channels={channels} />
                <ChannelForm />
            </div>
        )
    }
}

ReactDOM.render(<ChannelSection />, document.getElementById('app'));
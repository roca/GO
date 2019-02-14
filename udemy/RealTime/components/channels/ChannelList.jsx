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
import React from 'react';
import { Terminal } from 'xterm';
import './xterm.css';
import { FitAddon } from 'xterm-addon-fit'; 
import { WebLinksAddon } from 'xterm-addon-web-links';

class XTerminal extends React.Component {
    constructor(props) {
        super(props);
        this.terminal = new Terminal({cursorBlink: true});
        this.fitAddon = new FitAddon();
        const webLinksAddon = new WebLinksAddon();
        this.socket = new WebSocket('ws://localhost:8081/terminal');
        this.socket.onmessage=this.onSocketMessage;
        this.terminal.loadAddon(this.fitAddon);
        this.terminal.loadAddon(webLinksAddon);
        this.state = {
            response : ""
        }
    }

    componentDidMount() {
        const {id} = this.props;
        const terminalContainer = document.getElementById(id);
        this.terminal.open(terminalContainer);
        this.terminal.focus();
        this.terminal.onKey(this.onTerminalInput);
        this.fitAddon.fit();
    }

    onSocketMessage = (e) => {
        this.setState({response: e.data});
    };

    onTerminalInput = (e) => {
        this.socket.send(e.domEvent.key);
        if (e.domEvent.key === 'Enter') {
            this.terminal.write('\r\n');
        } else {
            this.terminal.write(e.key);
        }
    };

    render() {
        return(
            <div>
                <div id={this.props.id}/>
                <p>The Key You Enter:</p>
                <div>{this.state.response}</div>
            </div>
        )
    }
}

export default XTerminal;
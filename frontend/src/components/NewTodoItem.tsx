import React from "react";

import TextField from "@material-ui/core/TextField";
import IconButton from "@material-ui/core/IconButton";
import AddIcon from "@material-ui/icons/Add";

interface NewTodoItemProps {
}

interface NewTodoItemState {
    title: string;
}

class NewTodoItem extends React.Component<NewTodoItemProps, NewTodoItemState> {
    constructor(props: NewTodoItemProps) {
        super(props);

        this.state = {
            title: ""
        };
    }

    setTitle(title: string) {
        this.setState({
            title: title,
        });
    }

    render() {
        return (
            <div>
                <TextField
                    id="title"
                    label="Title"
                    value={this.state.title}
                    onChange={e => this.setTitle(e.target.value)}
                />
                <IconButton>
                    <AddIcon />
                </IconButton>
            </div>
        )
    }
}

export default NewTodoItem;

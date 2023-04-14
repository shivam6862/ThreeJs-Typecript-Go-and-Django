import { Component, ReactNode } from "react";

type CounterProps = {
  message: string;
};
type CounterState = {
  count: number;
};

// export class Couter extends Component<{}, CounterState> { // if no props
export class Couter extends Component<CounterProps, CounterState> {
  state = {
    count: 0,
  };

  handleClick = () => {
    this.setState((prevState) => ({ count: prevState.count + 1 }));
  };

  render() {
    return (
      <div>
        <h1>Class based components</h1>
        <div>
          <button onClick={this.handleClick}>Increment</button>
          <div>
            {this.props.message}
            {this.state.count}
          </div>
        </div>
      </div>
    );
  }
}

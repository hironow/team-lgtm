import React, { useState } from "react";
import * as ReactDOM from "react-dom";
import "normalize.css";
import "@blueprintjs/core/lib/css/blueprint.css";

import { Button } from "@blueprintjs/core";

// HelloProps
type HelloProps = {
  name: string;
};

// Hello
const Hello: React.FC<HelloProps> = ({ name }) => <h1>Hello {name}</h1>;

// App
const App: React.FC = () => {
  const [count, setCount] = useState(0);
  return (
    <React.Fragment>
      <Hello name="react" />
      <h3>
        Count: {count}
        <Button
          intent="success"
          text="Count"
          onClick={() => setCount(count + 1)}
        />
      </h3>
    </React.Fragment>
  );
};

ReactDOM.render(<App />, document.getElementById("root"));

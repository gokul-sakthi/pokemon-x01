import "./App.css";

import { Button } from "@/components/ui/button";
import { HStack } from "@chakra-ui/react";

const Demo = () => {
  return (
    <HStack>
      <Button>Click me</Button>
      <Button>Click me</Button>
    </HStack>
  );
};

function App() {
  return (
    <div className="App">
      <h1 className="bg-green-400">hello</h1>
      <Demo />
    </div>
  );
}

export default App;

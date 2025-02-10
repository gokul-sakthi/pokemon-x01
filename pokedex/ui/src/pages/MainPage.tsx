import { Box, Container, Input } from "@chakra-ui/react";
import { useState } from "react";
import PokemonCard from "./PokemonCard";

export const Results = () => {
  const [results] = useState<string[]>(["Pikachu"]);

  return (
    <Box>
      {results.map((result) => (
        <PokemonCard key={result} />
      ))}
    </Box>
  );
};

const MainPage = () => {
  return (
    <Container p={6}>
      <h1>Pokemon [ Pokedex ]</h1>

      <Input placeholder="Search" />
      <Results />
    </Container>
  );
};

export default MainPage;

import { Box, Button, Flex, Image, Span } from "@chakra-ui/react";

const PokemonCard = () => {
  return (
    <Flex
      w={96}
      borderRadius={10}
      minH={96}
      p={4}
      m={2}
      borderWidth={1}
      shadow={"md"}
      shadowColor="gray.400"
      borderColor="gray.400"
      direction={"column"}
    >
      <Flex
        justify="center"
        align="center"
        direction={"column"}
        bg="green.200"
        borderRadius={10}
        p={4}
        gap={4}
      >
        <Image w={200} h={200} src="/jigglipuff/33.png" alt="jigglypuff" />
      </Flex>

      <h3>Jigglypuff</h3>
      <Button> View </Button>
    </Flex>
  );
};

export default PokemonCard;

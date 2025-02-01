import { Flex, Heading, Text } from "@chakra-ui/react";
import { Switch } from "@/components/ui/switch";

const Navbar = () => {
  return (
    <Flex margin="4" justify="space-between">
      <Heading size="3xl">Pokedex</Heading>
      <Flex alignItems="center" gap="2">
        <Text> Dark Mode </Text>
        <Switch />
      </Flex>
    </Flex>
  );
};

export default Navbar;

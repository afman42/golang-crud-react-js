import React from "react";
import { Box, VStack } from "@chakra-ui/react";

const ViewList = () => {
  return (
    <div>
      <VStack p={4}>
        <Box bg="tomato" w="100%" p={4} color="white">
          This is the Box
        </Box>
      </VStack>
    </div>
  );
};

export default ViewList;

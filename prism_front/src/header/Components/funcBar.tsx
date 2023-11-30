import styled from "styled-components";
import { FaRegBell } from "react-icons/fa";
import { IoChatbubbleEllipsesOutline } from "react-icons/io5";
import { FaRegEdit } from "react-icons/fa";
import { BsPersonCircle } from "react-icons/bs";

const FuncBar: React.FC = () => {
  return (
    <Box>
      <ButtomBox>
        <FaRegBell size={"100%"} />
      </ButtomBox>
      <ButtomBox>
        <IoChatbubbleEllipsesOutline size={"100%"} />
      </ButtomBox>
      <ButtomBox>
        <FaRegEdit size={"100%"} />
      </ButtomBox>
      <ButtomBox>
        <BsPersonCircle size={"100%"} />
      </ButtomBox>
    </Box>
  );
};

export default FuncBar;

const Box = styled.div`
  display: flex;
  justify-content: right;
  height: 100%;
  align-items: center;
`;
const ButtomBox = styled.div`
  display: flex;
  align-items: center;
  margin: 0px 10px;
  cursor: pointer;
  height: 60%;
`;

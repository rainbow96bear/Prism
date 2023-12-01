import styled from "styled-components";
import { FaRegBell } from "react-icons/fa";
import { IoChatbubbleEllipsesOutline } from "react-icons/io5";
import { FaRegEdit } from "react-icons/fa";
import { BsPersonCircle } from "react-icons/bs";

const BeforeLogin: React.FC = () => {
  return (
    <>
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
    </>
  );
};

export default BeforeLogin;

const ButtomBox = styled.div`
  height: 50%;
  margin: 0px 10px;
  cursor: pointer;
`;

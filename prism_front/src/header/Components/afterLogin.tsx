import styled from "styled-components";
import { FaRegBell } from "react-icons/fa";
import { IoChatbubbleEllipsesOutline } from "react-icons/io5";
import { FaRegEdit } from "react-icons/fa";
import { useState } from "react";
import axios from "./../../configs/AxiosConfig";

import DropDown from "../../CustomComponent/DropDown";
import { TitlePath } from "../../GlobalType/TitlePath";

interface AfterLoginProps {
  userID?: string;
  imgUrl?: string;
}

const AfterLogin: React.FC<AfterLoginProps> = ({ userID, imgUrl }) => {
  const [dropdown, setDropdown] = useState(false);
  const logout = async () => {
    try {
      await axios.post("/OAuth/kakao/logout", {
        withCredentials: true,
      });
      const User_Login = document.cookie;
      const hasCookie = User_Login.includes("user_login");
      if (!hasCookie) {
        window.location.href = "http://localhost:3000";
      }
    } catch (error) {
      console.error("Logout failed:", error);
    }
  };
  const DropDown_List: TitlePath[] = [
    { title: "프로필", path: `/profile/${userID}` },
    { title: "로그아웃", func: logout },
  ];
  return (
    <>
      <ButtonBox>
        <FaRegBell size={"100%"} />
      </ButtonBox>
      <ButtonBox>
        <IoChatbubbleEllipsesOutline size={"100%"} />
      </ButtonBox>
      <ButtonBox>
        <FaRegEdit size={"100%"} />
      </ButtonBox>
      <ButtonBox>
        <ProfileImg
          onClick={() => setDropdown(!dropdown)}
          src={imgUrl}
          alt="User Profile"
        />

        {dropdown ? (
          <S_DropDown onClick={() => setDropdown(!dropdown)}>
            <DropDown list={DropDown_List} setDropdown={setDropdown} />
          </S_DropDown>
        ) : (
          <></>
        )}
      </ButtonBox>
    </>
  );
};

export default AfterLogin;

const ButtonBox = styled.div`
  position: relative;
  height: 50%;
  margin: 0px 10px;
  cursor: pointer;
`;

const ProfileImg = styled.img`
  height: 100%;
`;

const S_DropDown = styled.div`
  width: 100px;
  position: absolute;
  right: 0px;
  top: 50px;
  margin: 10px 0px;
`;

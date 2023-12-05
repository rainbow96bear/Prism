import styled from "styled-components";
import { FaRegBell } from "react-icons/fa";
import { IoChatbubbleEllipsesOutline } from "react-icons/io5";
import { FaRegEdit } from "react-icons/fa";
import { useEffect, useState } from "react";

import DropDown from "../../CustomComponent/DropDown";
import { TitlePath } from "../../GlobalType/TitlePath";

const BeforeLogin: React.FC = () => {
  const [imgSrc, setImgSrc] = useState("");
  const [profilePath, setProfilePath] = useState("");
  const [dropdown, setDropdown] = useState(false);

  useEffect(() => {
    // 쿠키에서 Img 및 ID 값 가져오기
    const kakaoUserValue = document.cookie;
    const newstring = kakaoUserValue.substring(1, kakaoUserValue.length - 1);
    const imgMatch = newstring.match(/Img=([^,]+)/);
    const imgValue = imgMatch ? decodeURIComponent(imgMatch[1]) : "";

    const IdMatch = newstring.match(/ID=([^,]+)/);
    const userIDValue = IdMatch ? decodeURIComponent(IdMatch[1]) : "";

    // Img의 값이 있다면 state에 저장
    if (imgValue) {
      setImgSrc(imgValue);
    }

    // ID 값이 있다면 profilePath를 생성하여 state에 저장
    if (userIDValue) {
      setProfilePath(`/profile/${userIDValue}`);
    }
  }, []);

  const DropDown_List: TitlePath[] = [{ title: "프로필", path: profilePath }];

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
          src={imgSrc}
          alt="User Profile"
        />

        {dropdown ? (
          <S_DropDown onClick={() => setDropdown(!dropdown)}>
            <DropDown list={DropDown_List} setDropdown={setDropdown}></DropDown>
          </S_DropDown>
        ) : (
          <></>
        )}
      </ButtonBox>
    </>
  );
};

export default BeforeLogin;

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

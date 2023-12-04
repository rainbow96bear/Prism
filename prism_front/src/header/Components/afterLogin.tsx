import styled from "styled-components";
import { FaRegBell } from "react-icons/fa";
import { IoChatbubbleEllipsesOutline } from "react-icons/io5";
import { FaRegEdit } from "react-icons/fa";
import { BsPersonCircle } from "react-icons/bs";
import { useEffect, useState } from "react";

const BeforeLogin: React.FC = () => {
  const [imgSrc, setImgSrc] = useState("");

  useEffect(() => {
    // 쿠키에서 Img의 값 가져오기
    const kakaoUserValue = document.cookie;
    const newstring = kakaoUserValue.substring(1, kakaoUserValue.length - 1);
    const imgMatch = newstring.match(/Img=([^,]+)/);
    const imgValue = imgMatch ? decodeURIComponent(imgMatch[1]) : "";

    // Img의 값이 있다면 state에 저장
    if (imgValue) {
      setImgSrc(imgValue);
    }
  }, []);
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
        {imgSrc == "" ? (
          <BsPersonCircle size={"100%"} />
        ) : (
          <ProfileImg src={imgSrc} alt="User Profile" />
        )}
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
const ProfileImg = styled.img`
  height: 100%;
`;

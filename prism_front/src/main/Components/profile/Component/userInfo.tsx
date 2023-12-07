import { useNavigate, useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import styled from "styled-components";
import axios from "axios";

const UserInfo = () => {
  const [imgSrc, setImgSrc] = useState("");
  const [nickName, setNickName] = useState("");
  const { id } = useParams();
  const [isUser, setIsUser] = useState(false);

  const navigator = useNavigate();

  const move = (path: string) => {
    navigator(path);
  };

  useEffect(() => {
    const kakaoUserValue = document.cookie;
    const newstring = kakaoUserValue.substring(1, kakaoUserValue.length - 1);
    const imgMatch = newstring.match(/Img=([^,]+)/);
    const imgValue = imgMatch ? decodeURIComponent(imgMatch[1]) : "";

    if (imgValue) {
      setImgSrc(imgValue);
    }

    const nickNameMatch = newstring.match(/NickName=([^,]+)/);
    const nickNameValue = nickNameMatch
      ? decodeURIComponent(nickNameMatch[1])
      : "";

    if (nickNameValue) {
      setNickName(nickNameValue);
    }

    const idMatch = newstring.match(/ID=([^,]+)/);
    const idValue = idMatch ? decodeURIComponent(idMatch[1]) : "";

    setIsUser(idValue === id);

    const getProfile = async () => {
      // const result = await axios.get(`http://localhost:8000/profile?id=${id}`);
    };
    getProfile();
  }, [id]);

  return (
    <Container>
      <img src={imgSrc} alt="User Profile" />
      <SubBox>
        <div className="nickname">{nickName}</div>
        <div className="oneLineIntroduce">한 줄 소개</div>
        <div className="hashtag">hashtag</div>
      </SubBox>
      <div>
        {isUser ? (
          <button
            onClick={() => {
              move("/profile/update/" + id);
            }}>
            수정
          </button>
        ) : (
          <button>메세지</button>
        )}
      </div>
    </Container>
  );
};

export default UserInfo;

const Container = styled.div`
  display: flex;
  height: 200px;

  > img {
    border-radius: 100%;
    width: 200px;
  }
`;

const SubBox = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  flex: 1;
  padding: 0px 20px;
  .nickname {
    font-weight: bold;
    font-size: 2rem;
  }
  .oneLineIntroduce {
    font-size: 1.2rem;
  }
  .hashtag {
    font-size: 1.2rem;
  }
`;

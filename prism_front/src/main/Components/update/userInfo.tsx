import styled from "styled-components";
import { useDispatch, useSelector } from "react-redux";
import { useEffect, useState } from "react";

import axios from "./../../../configs/AxiosConfig";
import { AppDispatch, RootState } from "../../../app/store";
import { getPersonalDate } from "../../../app/slices/profile/personal_data";
import { useNavigate } from "react-router-dom";
import HashTagItem from "./Component/hashtagItem";
import HashTagInput from "./Component/hashtagInput";

const UserInfo = () => {
  const dispatch = useDispatch<AppDispatch>();
  const user = useSelector((state: RootState) => state.user);
  const personalDate = useSelector((state: RootState) => state.personal_data);
  const navigator = useNavigate();

  const [nickname, setNickname] = useState(personalDate.nickname);
  const [hashTag, setHashTag] = useState(personalDate.hashtag);
  const [one_line_introduce, setOneLineIntroduce] = useState(
    personalDate.one_line_introduce
  );
  const [uploadedImage, setUploadedImage] = useState<string | null>(null);

  useEffect(() => {
    dispatch(getPersonalDate(user.user_id));
  }, []);

  const handleNicknameChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setNickname(event.target.value);
  };

  const handleIntroduceChange = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    setOneLineIntroduce(event.target.value);
  };

  const handleImageClick = () => {
    document.getElementById("imageInput")?.click();
  };

  const handleImageChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    if (file) {
      // 선택한 파일을 미리보기로 설정
      const reader = new FileReader();
      reader.onloadend = () => {
        setUploadedImage(reader.result as string);
      };
      reader.readAsDataURL(file);
    }
  };
  const handleHashTagRemove = (index: number) => {
    setHashTag((prevHashTag) => {
      const newHashTag = [...prevHashTag];
      newHashTag.splice(index, 1);
      return newHashTag;
    });
  };

  const saveChangedUserInfo = async () => {
    try {
      const userData = {
        nickname: nickname,
        one_line_introduce: one_line_introduce,
        profile_img: uploadedImage || personalDate.profile_img,
        hashtag: hashTag,
      };

      const response = await axios.post("/api/user/update", userData);

      if (response.status === 200) {
        console.log("User information updated successfully!");
        navigator(`/profile/${user.user_id}`);
      } else {
        console.error("Failed to update user information");
      }
    } catch (error) {
      console.error("Error during Axios request:", error);
    }
  };
  return (
    <Container>
      <InputBox>
        <img
          src={uploadedImage || personalDate.profile_img}
          alt="User Profile"
          onClick={handleImageClick}
          style={{ cursor: "pointer" }}
        />
        <input
          id="imageInput"
          type="file"
          accept="image/*"
          style={{ display: "none" }}
          onChange={handleImageChange}
        />
        <input
          type="text"
          value={nickname}
          onChange={handleNicknameChange}
          placeholder="닉네임"
        />
        <input
          type="text"
          value={one_line_introduce}
          onChange={handleIntroduceChange}
          placeholder="한 줄 소개"
        />
        <HashTagBox>
          {hashTag.map((value, index) => (
            <HashTagItem
              key={index}
              content={value}
              onRemove={() => handleHashTagRemove(index)}
            />
          ))}
          {hashTag.length < 5 && (
            <HashTagInput prevHashTag={hashTag} setHashTag={setHashTag} />
          )}
        </HashTagBox>
      </InputBox>
      <FuncBox>
        <div
          onClick={() => {
            navigator(`/profile/${user.user_id}`);
          }}>
          취소
        </div>
        <div
          onClick={() => {
            saveChangedUserInfo();
          }}>
          저장
        </div>
      </FuncBox>
    </Container>
  );
};

export default UserInfo;

const Container = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 50px;
`;

const InputBox = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  > img {
    width: 250px;
    padding: 0px 0px 50px 0px;
  }
  > input {
    width: 400px;
    padding: 10px;
    margin: 10px;
  }
`;

const HashTagBox = styled.div`
  display: flex;
  word-wrap: break-word;
`;

const FuncBox = styled.div`
  display: flex;
  justify-content: center;
  > div {
    border: solid 1px black;
    border-radius: 5px;
    margin: 10px;
    cursor: pointer;
    padding: 10px;
  }
`;

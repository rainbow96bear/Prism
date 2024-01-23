import styled from "styled-components";
import { useDispatch, useSelector } from "react-redux";
import { useEffect, useState } from "react";

import axios from "./../../../configs/AxiosConfig";
import { AppDispatch, RootState } from "../../../app/store";
import { getPersonalDate } from "../../../app/slices/profile/personal_data";
import { useNavigate } from "react-router-dom";
import HashTagItem from "./HashtagComponent/hashtagItem";
import HashTagInput from "./HashtagComponent/hashtagInput";
import ProfileImage from "../../../CustomComponent/ProfileImg";
import { fetchUser } from "../../../app/slices/user/user";

const UserInfo = () => {
  const dispatch = useDispatch<AppDispatch>();
  const user = useSelector((state: RootState) => state.user);
  const personalDate = useSelector((state: RootState) => state.personal_data);
  const navigator = useNavigate();
  const [nickname, setNickname] = useState<string | undefined>(undefined);
  const [hashTag, setHashTag] = useState(personalDate.hashtag);
  const [one_line_introduce, setOneLineIntroduce] = useState<
    string | undefined
  >(undefined);
  const [uploadedImage, setUploadedImage] = useState<string | null>(null);

  useEffect(() => {
    if (user.user_id) {
      dispatch(getPersonalDate(user.user_id));
      setOneLineIntroduce(personalDate.one_line_introduce);
      setNickname(personalDate.nickname);
    }
  }, [dispatch, user.user_id]);

  useEffect(() => {
    if (personalDate.nickname) {
      setOneLineIntroduce(personalDate.one_line_introduce);
      setNickname(personalDate.nickname);
    }
  }, [dispatch, personalDate.nickname]);

  useEffect(() => {
    dispatch(fetchUser());
  }, [dispatch]);

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
      const formData = new FormData();

      if (uploadedImage) {
        const imageFile = dataURItoBlob(uploadedImage);
        formData.append("image", imageFile, imageFile.type);
      }

      if (nickname != undefined) {
        formData.append("nickname", nickname);
      }
      if (one_line_introduce != undefined) {
        formData.append("one_line_introduce", one_line_introduce);
      }
      const hashtagsJsonString = JSON.stringify(hashTag);
      formData.append("hashtags", hashtagsJsonString);

      const response = await axios.post(
        `/profile/update/${user.user_id}`,
        formData,
        {
          withCredentials: true,
          headers: {
            "Content-Type": "multipart/form-data",
          },
        }
      );

      if (response.status === 200) {
        window.location.href = `/profile/${user.user_id}`;
      } else {
        console.error("Failed to update user information");
      }
    } catch (error) {
      console.error("Axios request error:", error);
    }
  };

  const dataURItoBlob = (dataURI: string) => {
    const byteString = atob(dataURI.split(",")[1]);
    const mimeString = dataURI.split(",")[0].split(":")[1].split(";")[0];
    const ab = new ArrayBuffer(byteString.length);
    const ia = new Uint8Array(ab);

    for (let i = 0; i < byteString.length; i++) {
      ia[i] = byteString.charCodeAt(i);
    }

    return new Blob([ab], { type: mimeString });
  };

  return (
    <Container>
      <InputBox>
        <ImageBox onClick={handleImageClick}>
          {uploadedImage != null ? (
            <img src={uploadedImage} alt="User Profile" />
          ) : (
            <ProfileImage id={user.user_id || "default"} />
          )}
        </ImageBox>
        <input
          id="imageInput"
          type="file"
          accept="image/*"
          style={{ display: "none" }}
          onChange={handleImageChange}
        />
        <Input
          type="text"
          value={nickname == undefined ? personalDate.nickname : nickname}
          onChange={handleNicknameChange}
          placeholder="닉네임"
        />
        <Input
          type="text"
          value={
            one_line_introduce == undefined
              ? personalDate.one_line_introduce
              : one_line_introduce
          }
          onChange={handleIntroduceChange}
          placeholder="한 줄 소개"
        />
        <HashTagBox>
          {hashTag?.map((value, index) => (
            <HashTagItem
              key={index}
              content={value}
              onRemove={() => handleHashTagRemove(index)}
            />
          ))}
          {hashTag == undefined && (
            <HashTagInput prevHashTag={hashTag} setHashTag={setHashTag} />
          )}
          {hashTag?.length < 5 && (
            <HashTagInput prevHashTag={hashTag} setHashTag={setHashTag} />
          )}
        </HashTagBox>
      </InputBox>
      <FuncBox>
        <Button onClick={() => navigator(`/profile/${user.user_id}`)}>
          취소
        </Button>
        <Button onClick={saveChangedUserInfo}>저장</Button>
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

const ImageBox = styled.div`
  width: 50%;
  position: relative;

  &:before {
    content: "";
    display: block;
    padding-bottom: 100%;
  }

  > img {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
`;

const Input = styled.input`
  width: 400px;
  padding: 10px;
  margin: 10px;
`;

const Button = styled.div`
  cursor: pointer;
  padding: 10px;
  background-color: #3498db;
  color: #fff;
  border: 1px solid #3498db;
  border-radius: 5px;
  margin-right: 10px;
  display: inline-block;

  &:last-child {
    margin-right: 0;
  }
`;

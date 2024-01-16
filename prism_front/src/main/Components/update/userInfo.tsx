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
  const [image, setImage] = useState<string | null>(null);

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
      // 이미지 업로드 요청
      const imageFormData = new FormData();
      if (uploadedImage) {
        const imageFile = dataURItoBlob(uploadedImage);
        imageFormData.append("file", imageFile, "profile_image.png");

        const imageResponse = await axios.post(
          "/profile/upload/image",
          imageFormData,
          {
            headers: {
              "Content-Type": "multipart/form-data",
            },
            withCredentials: true,
          }
        );
        console.log(imageResponse.data.imagePath);
        if (imageResponse.status === 200) {
          // 이미지 경로를 응답받은 경우에만 변경된 이미지 경로를 설정
          setImage(imageResponse.data.imagePath);
        }
      }

      // 프로필 정보 업데이트 요청
      const changedData = {
        nickname: nickname !== personalDate.nickname ? nickname : undefined,
        one_line_introduce:
          one_line_introduce !== personalDate.one_line_introduce
            ? one_line_introduce
            : undefined,
        profile_img: image !== personalDate.profile_img ? image : undefined,
        hashtag: hashTag,
      };

      // 변경된 값이 undefined인 경우 해당 프로퍼티 제거
      const userData = Object.fromEntries(
        Object.entries(changedData).filter(([_, value]) => value !== undefined)
      );

      const response = await axios.post(
        `/profile/update/${user.user_id}`,
        userData,
        {
          withCredentials: true,
        }
      );

      if (response.status === 200) {
        console.log("사용자 정보가 성공적으로 업데이트되었습니다!");
        navigator(`/profile/${user.user_id}`);
      } else {
        console.error("사용자 정보 업데이트에 실패했습니다");
      }
    } catch (error) {
      console.error("Axios 요청 중 오류 발생:", error);
    }
  };

  // Data URI를 Blob 객체로 변환
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

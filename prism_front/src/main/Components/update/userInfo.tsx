import styled from "styled-components";
import { useDispatch, useSelector } from "react-redux";
import { useEffect, useState } from "react";

import axios from "./../../../configs/AxiosConfig";
import { AppDispatch, RootState } from "../../../app/store";
import { getPersonalDate } from "../../../app/slices/profile/personal_data";
import { useNavigate } from "react-router-dom";
import HashTagItem from "./Component/hashtagItem";
import HashTagInput from "./Component/hashtagInput";
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
    // user_id가 존재할 때만 사용자 정보를 가져오도록 함
    if (user.user_id) {
      dispatch(getPersonalDate(user.user_id));
      setOneLineIntroduce(personalDate.one_line_introduce);
      setNickname(personalDate.nickname);
    }
  }, [dispatch, user.user_id]);

  useEffect(() => {
    // user_id가 존재할 때만 사용자 정보를 가져오도록 함
    if (personalDate.nickname) {
      setOneLineIntroduce(personalDate.one_line_introduce);
      setNickname(personalDate.nickname);
    }
  }, [dispatch, personalDate.nickname]);

  useEffect(() => {
    // 컴포넌트가 처음 마운트될 때에만 사용자 정보를 가져오도록 함
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
      const formData = new FormData();

      // Add image file to FormData
      if (uploadedImage) {
        const imageFile = dataURItoBlob(uploadedImage);
        formData.append("image", imageFile, imageFile.type);
      }

      // Add other form data
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
            "Content-Type": "multipart/form-data", // Set content type to multipart/form-data
          },
        }
      );

      if (response.status === 200) {
        window.location.href = `/profile/${user.user_id}`;
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
        <ImageBox onClick={handleImageClick} style={{ cursor: "pointer" }}>
          {uploadedImage != null ? (
            <img src={uploadedImage} alt="User Profile" />
          ) : (
            <ProfileImage
              id={user.user_id != "" ? user.user_id : "default"}></ProfileImage>
          )}
        </ImageBox>
        <input
          id="imageInput"
          type="file"
          accept="image/*"
          style={{ display: "none" }}
          onChange={handleImageChange}
        />
        <input
          type="text"
          value={nickname == undefined ? personalDate.nickname : nickname}
          onChange={handleNicknameChange}
          placeholder="닉네임"
        />
        <input
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

const ImageBox = styled.div`
  width: 50%;
  position: relative;

  &:before {
    content: "";
    display: block;
    padding-bottom: 100%; // 1:1 비율을 위한 값
  }

  > img {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    object-fit: cover; // 이미지가 비율을 유지하며 박스를 채우도록 함
  }
`;

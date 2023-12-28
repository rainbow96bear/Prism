// AddTechModal.tsx

import React, { useState } from "react";
import Modal from "react-modal";
import styled from "styled-components";
import axios from "axios";

interface AddTechModalProps {
  isOpen: boolean;
  onRequestClose: () => void;
  getTechList: () => void;
}

interface NewEntry {
  Tech_code: string;
  Tech_name: string;
}

const AddTechModal: React.FC<AddTechModalProps> = ({
  isOpen,
  onRequestClose,
  getTechList,
}) => {
  const [newEntry, setNewEntry] = useState<NewEntry>({
    Tech_code: "",
    Tech_name: "",
  });

  const handleAdd = async () => {
    try {
      // Send a request to add a new entry
      const result = (
        await axios.post("http://localhost:8080/admin/access/tech", newEntry, {
          withCredentials: true,
        })
      ).data;
      alert(
        "코드 번호 : " +
          result?.Tech_code +
          ", 기술명 : " +
          result?.Tech_name +
          "저장 완료"
      );
      getTechList();
      // Close the modal after adding successfully
      onRequestClose();
    } catch (error) {
      alert("중복된 코드를 입력하셨습니다.");
      console.error("Add failed:", error);
      // Handle error
    }
  };

  return (
    <Box>
      <Modal
        isOpen={isOpen}
        onRequestClose={onRequestClose}
        contentLabel="Add New Entry">
        <ModalContent>
          <h2>새로운 항목 추가</h2>
          <label>
            Tech_code:
            <input
              type="text"
              value={newEntry.Tech_code}
              onChange={(e) =>
                setNewEntry({ ...newEntry, Tech_code: e.target.value })
              }
            />
          </label>
          <label>
            Tech_name:
            <input
              type="text"
              value={newEntry.Tech_name}
              onChange={(e) =>
                setNewEntry({ ...newEntry, Tech_name: e.target.value })
              }
            />
          </label>
          <ButtonWrapper>
            <button onClick={handleAdd}>추가</button>
            <button onClick={onRequestClose}>취소</button>
          </ButtonWrapper>
        </ModalContent>
      </Modal>
    </Box>
  );
};

export default AddTechModal;

const Box = styled.div`
  top: "50%";
  left: "50%";
  right: "auto";
  bottom: "auto";
  marginright: "-50%";
  transform: "translate(-50%, -50%)";
`;

const ModalContent = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px;

  label {
    margin-top: 10px;
  }

  input {
    margin-left: 10px;
  }
`;

const ButtonWrapper = styled.div`
  margin-top: 20px;

  button {
    margin-right: 10px;
  }
`;

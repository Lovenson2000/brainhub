"use client";
import React, { useState } from "react";
import TaskCategory from "./TaskCategory";
import TaskDescription from "./TaskDescription";
import TaskPriority from "./TaskPriority";
import TaskStatus from "./TaskStatus";
import TaskTitle from "./TaskTitle";
import { DotsVerticalIcon } from "@heroicons/react/solid";
import TaskMenu from "./TaskMenuModal";

type TaskType = {
  title: string;
  description: string;
  category: string;
  dueDate: string;
  status: "To Do" | "In Progress" | "Done";
  priority: "Low" | "Normal" | "High";
};

const Task: TaskType = {
  title: "Calculus Review",
  description: "Review first three chapters of calculus II",
  category: "Homework",
  dueDate: "Wed, 23 Aug",
  status: "In Progress",
  priority: "Normal",
};

export default function TaskItem() {
  const [menuVisible, setMenuVisible] = useState(false);

  const handleEdit = () => {
    console.log("Edit Task");
    setMenuVisible(false);
  };

  const handleDelete = () => {
    console.log("Delete Task");
    setMenuVisible(false);
  };

  const handleFavorite = () => {
    console.log("Add to Favorite");
    setMenuVisible(false);
  };

  const handleMarkAsDone = () => {
    console.log("Mark as Done");
    setMenuVisible(false);
  };

  return (
    <div className="relative border max-w-[14rem] p-2 rounded shadow-sm">
      <div className="absolute top-2 right-2">
        <button onClick={() => setMenuVisible(!menuVisible)}>
          <DotsVerticalIcon className="h-5 w-5 text-gray-600" />
        </button>
        {menuVisible && (
          <TaskMenu
            onEdit={handleEdit}
            onDelete={handleDelete}
            onFavorite={handleFavorite}
            onMarkAsDone={handleMarkAsDone}
          />
        )}
      </div>
      <TaskCategory category={Task.category} />
      <TaskTitle title={Task.title} />
      <TaskDescription description={Task.description} />
      <div className="flex gap-4 m-1">
        <TaskStatus status={Task.status} />
        <TaskPriority priority={Task.priority} />
      </div>
    </div>
  );
}



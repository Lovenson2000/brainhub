import React from "react";

type TaskStatusProps = {
  status: "To Do" | "In Progress" | "Done";
};
export default function TaskStatus({ status }: TaskStatusProps) {
  return (
    <div className="font-light p-1 rounded-sm text-sm text-slate-900 border">
      <h3>{status}</h3>
    </div>
  );
}

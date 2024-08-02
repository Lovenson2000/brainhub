import React from "react";

type TaskPriorityProps = {
  priority: "Low" | "Normal" | "High";
};

export default function TaskPriority({ priority }: TaskPriorityProps) {
  return (
    <div className="p-1 font-light rounded-sm text-sm text-slate-900 border">
      <h3>{priority}</h3>
    </div>
  );
}

import React from "react";
import {
  PencilAltIcon,
  TrashIcon,
  StarIcon,
  CheckIcon,
} from "@heroicons/react/solid";

type TaskMenuProps = {
  onEdit: () => void;
  onDelete: () => void;
  onFavorite: () => void;
  onMarkAsDone: () => void;
};

export default function TaskMenu({
  onEdit,
  onDelete,
  onFavorite,
  onMarkAsDone,
}: TaskMenuProps) {
  return (
    <div className="absolute p-1 right-0 mt-2 w-48 bg-white border rounded shadow-lg z-10">
      <ul className="p-1 flex flex-col gap-2">
        <MenuItem
          Icon={PencilAltIcon}
          text="Edit Task"
          iconStyle=""
          onClick={onEdit}
        />
        <MenuItem
          Icon={TrashIcon}
          text="Delete Task"
          iconStyle=""
          onClick={onDelete}
        />
        <MenuItem
          Icon={StarIcon}
          text="Add to Favorite"
          iconStyle=""
          onClick={onFavorite}
        />
        <MenuItem
          Icon={CheckIcon}
          text="Mark as Done"
          iconStyle=""
          onClick={onMarkAsDone}
        />
      </ul>
    </div>
  );
}

type MenuItemProps = {
  Icon: React.ElementType;
  text: string;
  iconStyle: string;
  onClick: () => void;
};

function MenuItem({ Icon, text, iconStyle, onClick }: MenuItemProps) {
  return (
    <li
      className="flex items-center px-4 py-1 rounded text-gray-700 hover:bg-gray-100 cursor-pointer"
      onClick={onClick}
    >
      <Icon className={`h-5 w-5 mr-2 ${iconStyle}`} />
      {text}
    </li>
  );
}

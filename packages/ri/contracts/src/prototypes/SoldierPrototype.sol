// SPDX-License-Identifier: Unlicense
pragma solidity >=0.8.0;

import { IUint256Component } from "solecs/interfaces/IUint256Component.sol";
import { getAddressById } from "solecs/utils.sol";

import { LibPrototype } from "../libraries/LibPrototype.sol";

import { PrototypeComponent, ID as PrototypeComponentID } from "../components/PrototypeComponent.sol";
import { UnitTypeComponent, ID as UnitTypeComponentID } from "../components/UnitTypeComponent.sol";
import { MovableComponent, ID as MovableComponentID } from "../components/MovableComponent.sol";
import { StaminaComponent, Stamina, ID as StaminaComponentID } from "../components/StaminaComponent.sol";
import { CombatComponent, Combat, ID as CombatComponentID } from "../components/CombatComponent.sol";
import { UntraversableComponent, ID as UntraversableComponentID } from "../components/UntraversableComponent.sol";
import { InventoryComponent, ID as InventoryComponentID } from "../components/InventoryComponent.sol";

import { UnitTypes, CombatTypes } from "../utils/Types.sol";

uint256 constant ID = uint256(keccak256("mudwar.prototype.Soldier"));

function SoldierPrototype(IUint256Component components) {
  UnitTypeComponent(getAddressById(components, UnitTypeComponentID)).set(ID, uint32(UnitTypes.Soldier));
  StaminaComponent(getAddressById(components, StaminaComponentID)).set(
    ID,
    Stamina({ current: 0, max: 3, regeneration: 1 })
  );
  CombatComponent(getAddressById(components, CombatComponentID)).set(
    ID,
    Combat({ _type: uint32(CombatTypes.Sword), strength: 20_000, health: 100_000, passive: false })
  );
  MovableComponent(getAddressById(components, MovableComponentID)).set(ID, int32(4));
  UntraversableComponent(getAddressById(components, UntraversableComponentID)).set(ID);
  InventoryComponent(getAddressById(components, InventoryComponentID)).set(ID, uint32(3));

  uint256[] memory componentIds = new uint256[](6);
  componentIds[0] = UnitTypeComponentID;
  componentIds[1] = StaminaComponentID;
  componentIds[2] = MovableComponentID;
  componentIds[3] = CombatComponentID;
  componentIds[4] = UntraversableComponentID;
  componentIds[5] = InventoryComponentID;

  PrototypeComponent(getAddressById(components, PrototypeComponentID)).set(ID, componentIds);
}
